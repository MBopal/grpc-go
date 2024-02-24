package main

import (
	"context"
	"log"

	pb "github.com/PenghapusGwIlang/grpc-go/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---readBlog was invoked---")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
