package main

import (
	"context"
	"log"

	pb "github.com/erikrios/bookshop/client/pb/inventory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect: %w", err)
	}
	defer conn.Close()

	client := pb.NewInventoryClient(conn)

	bookList, err := client.GetBookList(context.Background(), &pb.GetBookListRequest{})

	log.Printf("book list: %v", bookList)
}
