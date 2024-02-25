package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/PenghapusGwIlang/grpc-go/blog/proto"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)          //valid
	readBlog(c, "Balabalbs") //invalid
}