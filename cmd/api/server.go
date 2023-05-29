package api

import (
	"context"
	"jet/config"
	"jet/ent"
	"jet/pb"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServerCmd(configs *config.Configurations, logger *zap.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "run api server",
		Long:  "run api server with graphql",
		Run: func(cmd *cobra.Command, args []string) {

			// Connect to postgresql database
			db, err := ent.Open("postgres", configs.Postgres.ConnectionString)
			if err != nil {
				logger.Error("Getting error connect to postgresql database", zap.Error(err))
				os.Exit(1)
			}
			defer db.Close()

			// Run the automation migration tool
			if err := db.Schema.Create(context.Background()); err != nil {
				logger.Error("Failed to creating db schema from the automation migration tool", zap.Error(err))
				os.Exit(1)
			}
			//Conenct to Customer server
			conn1, err := grpc.Dial(":1000", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("Failed connecting to Customer server: %s", err)
			}
			defer conn1.Close()
			// Create validator
			validator := validator.New()
			en := en.New()
			uni := ut.New(en, en)
			validationTranslator, _ := uni.GetTranslator("en")

			accountService := pb.NewAccountServiceClient(conn1)

			// Register default translation for validator
			err = en_translations.RegisterDefaultTranslations(validator, validationTranslator)
			if err != nil {
				logger.Error("Getting error from register default translation", zap.Error(err))
				os.Exit(1)
			}

			// GraphQL schema resolver handler.
			resolverHandler := handler.NewDefaultServer(resolver.NewSchema(db, validator, validationTranslator, logger, customerService, accountService, flightService, passengerService, ticketService, bookingService))
			// Handler for GraphQL Playground
			playgroundHandler := playground.Handler("GraphQL Playground", "/graphql")

			// Handle a method not allowed.
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.HandleMethodNotAllowed = true

			// Use middlewares
			r.Use(
				auth.Auth(accountService),
				ginzap.Ginzap(logger, time.RFC3339, true),
				ginzap.RecoveryWithZap(logger, true),
				middleware.CorsMiddleware(),
				middleware.RequestCtxMiddleware(),
			)

			// Create a new GraphQL
			r.POST("/graphql", func(c *gin.Context) {
				resolverHandler.ServeHTTP(c.Writer, c.Request)
			})

			r.OPTIONS("/graphql", func(c *gin.Context) {
				c.Status(200)
			})

			// Enable playground for query/testing
			r.GET("/", func(c *gin.Context) {
				playgroundHandler.ServeHTTP(c.Writer, c.Request)
			})

			// Listen on port 8000
			logger.Info("Listening on port: 8000")
			if err := r.Run(":8000"); err != nil {
				logger.Error("Get error from run server", zap.Error(err))
			}
		},
	}
}
