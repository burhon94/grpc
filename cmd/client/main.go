package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/pkg/api"
	"log"
)

func main() {
	var x, y int

	fmt.Print("x = ")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("y = ")
	_, err = fmt.Scanf("%d", &y)
	if err != nil {
		log.Fatal(err)
	}

	dial, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	newClient := api.NewAdderClient(dial)

	addResp, err := newClient.Add(context.Background(), &api.AddReq{
		X: int32(x),
		Y: int32(y),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("gRPC response: %v", addResp.Result)
}
