package main

import (
	"context"
	"log"
	"time"

	pb "github.com/PenghapusGwIlang/grpc-go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet function was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Naufal"},
		{FirstName: "Nopal"},
		{FirstName: "Bopal"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}

	log.Printf("LongGreet response: %v\n", res.Result)
}
