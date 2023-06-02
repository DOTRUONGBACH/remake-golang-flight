package auth

import (
	"context"
	"jet/pb"

	"jet/internal/jwt"

	"github.com/gin-gonic/gin"
)

type contextKey struct {
	name string
}

type accConText struct {
	AccEmail string
	Role     string
}

var userCtxKey = &contextKey{"user"}

func Auth(s pb.AccountServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		token := authHeader[7:]
		accountEmail, err := jwt.ParseToken(token)
		if err != nil {
			c.Next()
			return
		}
		r := c.Request
		ctx := r.Context()
		pbAcount, err := s.GetAccountByEmail(ctx, &pb.GetAccountByEmailRequest{Email: accountEmail})
		if err != nil {
			c.Next()
			return
		}

		acc := &accConText{AccEmail: accountEmail, Role: pbAcount.Role.String()}
		ctx = context.WithValue(ctx, userCtxKey, acc)
		c.Request = r.WithContext(ctx)

		c.Next()

	}
}
