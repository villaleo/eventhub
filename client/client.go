package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/villaleo/eventhub/eventhub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddr = flag.String(
	"addr",
	"localhost:50051",
	"The server address in the format of host:port",
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

	c := pb.NewEventManagerClient(rpcClient)
	event, err := c.NewEvent(context.Background(), &pb.Event{
		Name:        "SOMOS AI Workshop",
		Description: "It's all the way in Santa Cruz tho",
		Timestamp:   time.Now().Local().String(),
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(event)
}
