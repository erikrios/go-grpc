package main

import (
	"context"
	"log"
	"net"

	pb "github.com/erikrios/bookshop/server/pb/inventory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterInventoryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	return &pb.GetBookListResponse{
		Books: getSampleBook(),
	}, nil
}

func getSampleBook() []*pb.Book {
	sampleBooks := []*pb.Book{
		{
			Title:     "The Hitchhikker's Guide to the Galaxy",
			Author:    "Douglas Adams",
			PageCount: 42,
		},
		{
			Title:     "The Lord of the Rings",
			Author:    "J.R.R Tolkien",
			PageCount: 1234,
		},
	}
	return sampleBooks
}
