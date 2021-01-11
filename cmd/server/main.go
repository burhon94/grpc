package main

import (
	"grpc/pkg/adder"
	"grpc/pkg/api"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatal(err)
	}

	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}