package main

import (
	"context"
	"log"

	pb "github.com/PenghapusGwIlang/grpc-go/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg function was invoked")

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}

	numbers := []int32{1, 2, 3, 4, 5}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)

		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}
