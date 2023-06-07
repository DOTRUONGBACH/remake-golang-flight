package main

import (
	"jet/grpc/customer_grpc/handler"
	"jet/grpc/customer_grpc/repository"
	"jet/pb"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":2223")
	if err != nil {
		panic(err)
	}

	// logger, _ := zap.NewProduction()
	// defer logger.Sync()

	s := grpc.NewServer()

	customerRepository, err := repository.NewPostgresDB("host=localhost port=5432 user=postgres dbname=flight password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	// defer userRepository.Close()

	h, err := handler.NewCustomerrHandler(customerRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterCustomerServiceServer(s, h)

	// logger.Info("Listen at port: 2223")

	log.Println("CUSTOMER server is listening at port 2223...")
	s.Serve(listen)
}
