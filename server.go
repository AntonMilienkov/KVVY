package main

import (
	"context"
	"log"
	pb "nauchka/gRPC"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.GraphServiceServer
}

func (s *server) SendNode(ctx context.Context, in *pb.NodeRequest) (*pb.NodeResponse, error) {
	return &pb.NodeResponse{Message: "Hello, World! "}, nil
}

func serverStart() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGraphServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
