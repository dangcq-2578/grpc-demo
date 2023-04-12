package main

import (
	"context"
	"fmt"
	"log"
	"time"

	services "github.com/dangcq-2578/grpc-demo/proto"

	"google.golang.org/grpc"
)

var (
	severURL = "localhost:9123"
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

	posts, err := client.GetPosts(ctx, &services.Empty{})

	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts.GetPosts() {
		fmt.Println(post.Id)
		fmt.Println(post.Title)
		fmt.Println(post.Content)
	}
}
