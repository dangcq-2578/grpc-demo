package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

func getPostgRPC() *services.PostList {
	conn := getgRPCClient()

	defer conn.Close()

	client := services.NewPostServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	posts, err := client.GetPosts(ctx, &services.Empty{})

	if err != nil {
		log.Fatal(err)
	}

	return posts
}

func main() {
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		posts := getPostgRPC()

		for _, post := range posts.GetPosts() {
			fmt.Fprintln(w, post.Id)
			fmt.Fprintln(w, post.Title)
			fmt.Fprintln(w, post.Content)
		}
	})

	http.ListenAndServe(":8081", nil)
}
