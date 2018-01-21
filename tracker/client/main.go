package main

import (
	"log"

	"context"

	"github.com/darren-west/marathon/tracker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := tracker.NewLocationServiceClient(conn)

	md := metadata.New(map[string]string{"key1": "val1", "key2": "val2"})

	ctx := metadata.NewOutgoingContext(context.TODO(), md)

	client.AddLocation(ctx, &tracker.Location{})
}
