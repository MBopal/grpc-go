package main

import (
	"context"
	"log"

	pb "github.com/PenghapusGwIlang/grpc-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("DoGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Bopal",
	})

	if err != nil {
		log.Fatalf("Failed to greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
