package main

import (
	"context"
	"log"
	pb "nauchka/gRPC"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func client(data, hashSelf, hashOther, hashOwn string, ts int64) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()
	c := pb.NewGraphServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.WriteDataToAnotherNode(ctx, &pb.NodeRequest{Data: data, HashSelfParent: hashSelf, HashOtherParent: hashOther, HashOwn: hashOwn, Timestamp: ts})
	if err != nil {
		log.Fatalf("error calling function WriteDataToAnotherNode: %v", err)
	}

	log.Printf("Response from gRPC server's WriteDataToAnotherNode function: %s", r.GetMessage())
}
