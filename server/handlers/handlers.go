package handlers

import (
	"context"
	"fmt"

	pb "Exercise/server/proto"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

const (
	tableName = "ExerciseTableV1"
)

type Server struct {
	pb.UnimplementedExerciseServiceServer
	DBClient *dynamodb.Client
}

func (s *Server) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	item := map[string]types.AttributeValue{
		"id":    &types.AttributeValueMemberS{Value: req.Id},
		"name":  &types.AttributeValueMemberS{Value: req.Name},
		"value": &types.AttributeValueMemberN{Value: req.Value},
	}

	_, err := s.DBClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create item: %w", err)
	}

	return &pb.CreateItemResponse{Success: true}, nil
}

func (s *Server) ReadItem(ctx context.Context, req *pb.ReadItemRequest) (*pb.ReadItemResponse, error) {
	result, err := s.DBClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: req.Id},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read item: %w", err)
	}

	if result.Item == nil {
		return nil, fmt.Errorf("item not found")
	}

	nameAttr, nameOk := result.Item["name"].(*types.AttributeValueMemberS)
	valueAttr, valueOk := result.Item["value"].(*types.AttributeValueMemberN)
	if !nameOk || !valueOk {
		return nil, fmt.Errorf("invalid item structure")
	}

	return &pb.ReadItemResponse{
		Name:  nameAttr.Value,
		Value: valueAttr.Value,
	}, nil
}

func (s *Server) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	_, err := s.DBClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: req.Id},
		},
		UpdateExpression: aws.String("SET #name = :name, #value = :value"),
		ExpressionAttributeNames: map[string]string{
			"#name":  "name",
			"#value": "value",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name":  &types.AttributeValueMemberS{Value: req.Name},
			":value": &types.AttributeValueMemberN{Value: req.Value},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update item: %w", err)
	}

	return &pb.UpdateItemResponse{Success: true}, nil
}

func (s *Server) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	_, err := s.DBClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: req.Id},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to delete item: %w", err)
	}

	return &pb.DeleteItemResponse{Success: true}, nil
}
