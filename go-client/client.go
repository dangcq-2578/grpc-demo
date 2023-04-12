package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dangcq-2578/grpc-demo/services"

	"google.golang.org/grpc"
)

var (
	severURL = "localhost:9321"
)

func getgRPCClient() *grpc.ClientConn {
	var opts = []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	conn, err := grpc.Dial(severURL, opts...)
	if err != nil {
		log.Fatalf("Fail to dial: %v", err)
	}

	return conn
}

func main() {
	conn := getgRPCClient()

	defer conn.Close()

	client := services.NewPostServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	posts, err := client.GetPost(ctx, &services.Empty{})

	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts.getPosts() {
		fmt.Println(post.Id)
		fmt.Println(post.Title)
		fmt.Println(post.Text)
	}
}
