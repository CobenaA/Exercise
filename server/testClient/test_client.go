package main

import (
	"context"
	"log"
	"time"

	pb "Exercise/server/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewExerciseServiceClient(conn)

	// Create an Item
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createRes, err := client.CreateItem(ctx, &pb.CreateItemRequest{
		Id:    "1",
		Name:  "Test Item",
		Value: "100",
	})
	if err != nil {
		log.Fatalf("could not create item: %v", err)
	}
	log.Printf("CreateItem response: %v", createRes)

	// Read the Item
	readRes, err := client.ReadItem(ctx, &pb.ReadItemRequest{
		Id: "1",
	})
	if err != nil {
		log.Fatalf("could not read item: %v", err)
	}
	log.Printf("ReadItem response: %v", readRes)
}
