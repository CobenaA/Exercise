package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"google.golang.org/grpc"

	"Exercise/server/handlers"

	pb "Exercise/server/proto"
)

func main() {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("us-west-2"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				if service == dynamodb.ServiceID {
					return aws.Endpoint{URL: "http://localhost:8000"}, nil
				}
				return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
			}),
		),
	)
	if err != nil {
		log.Fatalf("failed to load AWS configuration: %v", err)
	}

	// Initialize DynamoDB client
	dbClient := dynamodb.NewFromConfig(cfg)

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register ExerciseService server implementation
	exerciseServer := &handlers.Server{DBClient: dbClient}
	pb.RegisterExerciseServiceServer(grpcServer, exerciseServer)

	// Start listening on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	log.Println("gRPC server is listening on port 50051")

	// Serve gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
