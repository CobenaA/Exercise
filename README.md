# Exercise Project

The **Exercise Project** is a Go-based backend that provides gRPC-based CRUD operations. It is designed to integrate with a local DynamoDB instance for development and testing. This repository sets the foundation for managing resources and will be extended to support `Exercise` models in future updates.

## Features
- Basic gRPC CRUD operations.
- Local DynamoDB integration for development.

## Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install) (1.20 or later recommended)
- [Protocol Buffers Compiler (protoc)](https://grpc.io/docs/protoc-installation/)
- [AWS DynamoDB Local](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.html)

### Setting Up Local DynamoDB
1. **Run DynamoDB Local with Docker**:
   ```bash
   docker run -p 8000:8000 amazon/dynamodb-local
   ```

2. **Verify DynamoDB Local is Running**:
   ```bash
   aws dynamodb list-tables --endpoint-url http://localhost:8000
   ```
   If successful, you'll see an empty list of tables initially.

3. **Create a Table**:
   ```bash
   aws dynamodb create-table \
       --table-name ExerciseTableV1 \
       --attribute-definitions AttributeName=id,AttributeType=S \
       --key-schema AttributeName=id,KeyType=HASH \
       --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
       --endpoint-url http://localhost:8000
   ```

### Running the Project
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/Exercise.git
   cd Exercise
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Start the gRPC server:
   ```bash
   go run server/main.go
   ```
   The server will start listening on `localhost:50051`.

### Testing the gRPC Server
You can use the included `server/testClient/test_client.go` file to test the CRUD operations. The client provides a straightforward way to interact with the server. To run the test client:

1. Navigate to the client directory:
   ```bash
   cd server/testClient
   ```

2. Run the test client:
   ```bash
   go run test_client.go
   ```

The client will execute a series of CRUD operations and display the results in the console.

## Future Plans
- Integrate `Exercise` models into the CRUD operations.
- Add advanced validation and error handling.
- Extend functionality for more complex operations.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

