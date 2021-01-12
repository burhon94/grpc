package main

import (
	"grpc/pkg/api"
	"grpc/pkg/gServer"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &gServer.GRPCServer{}
	api.RegisterGServiceServer(s, srv)

	l, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatal(err)
	}

	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}