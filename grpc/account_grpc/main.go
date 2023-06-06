package main

import (
	"jet/grpc/account_grpc/handler"
	"jet/grpc/account_grpc/repository"
	"jet/pb"
	"log"
	"net"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":2224")
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	s := grpc.NewServer()
	log.Println(s, listen)
	// accountrepo, err :=

	accountRepository, err := repository.NewPostgresDB("host=localhost port=5432 user=postgres dbname=flight password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	// defer repository.AccountRepository.CloseDB(accountRepository)

	h, err := handler.NewAccountHander(accountRepository)
	if err != nil {
		panic(err)
	}
	reflection.Register(s)
	pb.RegisterAccountServiceServer(s, h)

	log.Println("CUSTOMER server is listening at port 2224...")
	s.Serve(listen)

}
