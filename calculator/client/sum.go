package main

import (
	"context"
	"log"

	pb "github.com/PenghapusGwIlang/grpc-go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	})

	if err != nil {
		log.Fatalf("Failed to sum: %v", err)
	}

	log.Printf("sum: %d\n", res.Result)
}
