package main

import (
	"context"
	"log"

	pb "github.com/PenghapusGwIlang/grpc-go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog was invoked---")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Failed to delete blog: %v\n", err)
	}

	log.Println("Blog deleted")
}
