package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/pkg/api"
	"log"
	"os"
)

func main() {
	for {

		var x, y int

		var choose int
		fmt.Printf("\n\n----------\n" +
			"select clientMode: \n" +
			"1-Adder\n" +
			"2-Minuser\n" +
			"3-Multipler\n" +
			"4-Diveder\n" +
			"5-PercentAdder\n" +
			"6-PercentMinus\n" +
			"Yor Choose: ")
		_, err2 := fmt.Scanf("%d", &choose)
		if err2 != nil {
			log.Fatal("error data")

			return
		}
		if choose == 0 {
			os.Exit(0)
		}
		switch choose {
		case 1:
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

			newClient := api.NewGServiceClient(dial)

			addResp, err := newClient.Add(context.Background(), &api.Req{
				X: int32(x),
				Y: int32(y),
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("gRPC response: %v", addResp.Result)
		case 2:
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

			newClient := api.NewGServiceClient(dial)

			minusResp, err := newClient.Minus(context.Background(), &api.Req{
				X: int32(x),
				Y: int32(y),
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("gRPC response: %v", minusResp.Result)
		case 3:
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

			newClient := api.NewGServiceClient(dial)

			minusResp, err := newClient.Multiple(context.Background(), &api.Req{
				X: int32(x),
				Y: int32(y),
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("gRPC response: %v", minusResp.Result)
		case 4:
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

			newClient := api.NewGServiceClient(dial)

			minusResp, err := newClient.Dived(context.Background(), &api.Req{
				X: int32(x),
				Y: int32(y),
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("gRPC response: %v", minusResp.Result)
		case 5:
			fmt.Print("x(100%) + y(%) = z(%)\n")
			fmt.Print("x(100%) = ")
			_, err := fmt.Scanf("%d", &x)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Print("y(%) = ")
			_, err = fmt.Scanf("%d", &y)
			if err != nil {
				log.Fatal(err)
			}

			dial, err := grpc.Dial(":8080", grpc.WithInsecure())
			if err != nil {
				log.Fatal(err)
			}

			newClient := api.NewGServiceClient(dial)

			minusResp, err := newClient.PercentAdd(context.Background(), &api.Req{
				X: int32(x),
				Y: int32(y),
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("gRPC response: %v", minusResp.Result)
		case 6:
			fmt.Print("x(100%) - y(%) = z(%)\n")
			fmt.Print("x(100%) = ")
			_, err := fmt.Scanf("%d", &x)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Print("y(%) = ")
			_, err = fmt.Scanf("%d", &y)
			if err != nil {
				log.Fatal(err)
			}

			dial, err := grpc.Dial(":8080", grpc.WithInsecure())
			if err != nil {
				log.Fatal(err)
			}

			newClient := api.NewGServiceClient(dial)

			minusResp, err := newClient.PercentMinus(context.Background(), &api.Req{
				X: int32(x),
				Y: int32(y),
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("gRPC response: %v", minusResp.Result)

		default:

			log.Fatal("error data")

		}

	}

}
