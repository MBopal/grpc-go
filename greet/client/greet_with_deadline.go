package main

import (
	"context"
	"log"
	"time"

	pb "github.com/PenghapusGwIlang/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline function was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{FirstName: "Naufal"}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline was exceeded")
				return
			} else {
				log.Printf("Unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error: %v", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
