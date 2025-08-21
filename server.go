package main

import (
	"context"
	"log"
	"net"

	"nauchka/files"
	pb "nauchka/gRPC"
	"nauchka/graph"

	"google.golang.org/grpc"
)

type server struct {
	pb.GraphServiceServer
}

func (s *server) WriteDataToAnotherNode(ctx context.Context, in *pb.NodeRequest) (*pb.NodeResponse, error) {
	var node graph.Node

	node.Data = in.GetData()
	node.HashSelfParent = in.GetHashSelfParent()
	node.HashOtherParent = in.GetHashOtherParent()
	node.HashOwn = in.GetHashOwn()
	node.Timestamp = in.GetTimestamp()

	files.WriteToFile(node)

	return &pb.NodeResponse{Message: "Node recieved"}, nil
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
