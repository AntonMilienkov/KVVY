package graph

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"nauchka/files"
	pb "nauchka/gRPC"

	"google.golang.org/grpc"
)

type server struct {
	pb.GraphServiceServer
}

func (s *server) WriteDataToAnotherNode(ctx context.Context, in *pb.NodeRequest) (*pb.NodeResponse, error) {
	var node Node

	node.Data = in.GetData()
	node.HashSelfParent = in.GetHashSelfParent()
	node.HashOtherParent = in.GetHashOtherParent()
	node.HashOwn = in.GetHashOwn()
	node.Timestamp = in.GetTimestamp()

	files.WriteToFile(node)

	return &pb.NodeResponse{Message: "Node recieved"}, nil
}

func ServerStart() {
	port := strconv.Itoa(hosts[selfNumber].Port)
	fmt.Println("Server Start at " + port)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterGraphServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
