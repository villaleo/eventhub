// Package main implements a RouteGuide client.
package main

import (
	"flag"
	"fmt"
	"log"

	pb "github.com/villaleo/eventhub/eventhub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String(
		"addr",
		"localhost:50051",
		"The server address in the format of host:port",
	)
)

func main() {
	flag.Parse()

	rpcClient, err := grpc.NewClient(
		fmt.Sprintf("dns:///%s", *serverAddr),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("failed to dial:", err)
	}
	defer rpcClient.Close()

	_ = pb.NewEventManagerClient(rpcClient)
}
