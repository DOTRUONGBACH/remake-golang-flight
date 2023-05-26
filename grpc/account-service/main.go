package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "1000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	log.Println(s, listen)
	// accountrepo, err :=

}
