package graph

import (
	"context"
	"fmt"
	"log"
	pb "nauchka/gRPC"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func client(host Host, noda *Node) {
	address := host.DnsName + ":" + strconv.Itoa(host.Port)
	fmt.Println("client connect to " + address)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at %s: %v", address, err)
	}
	defer conn.Close()
	c := pb.NewGraphServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.WriteDataToAnotherNode(ctx, &pb.NodeRequest{
		Data:            noda.Data,
		HashSelfParent:  noda.HashSelfParent,
		HashOtherParent: noda.HashOtherParent,
		HashOwn:         noda.HashOwn,
		Timestamp:       noda.Timestamp,
	})

	if err != nil {
		log.Fatalf("error calling function WriteDataToAnotherNode: %v", err)
	}

	log.Printf("Response from gRPC server's WriteDataToAnotherNode function: %s", r.GetMessage())
}
