package main

import (
	"context"
	"log"

	pb "github.com/PenghapusGwIlang/grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Naufal",
		Title:    "New Title",
		Content:  "Content of the first with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Could not update: %v\n", err)
	}

	log.Println("Blog was updated")
}
