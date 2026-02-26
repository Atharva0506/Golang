package main

import (
	"context"
	"log"
	"log/slog"
	"net"

	pb "grpc_demo/proto" // This imports the incredible code gRPC built for us!

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 1. The Server Struct
// We MUST embed the Unimplemented server that gRPC generated for us to be safe!
type server struct {
	pb.UnimplementedCalculatorServiceServer
}

// 2. The Implementation
// gRPC forces this exact signature! We just write the business logic.
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	// Add req.A and req.B together, and return a *pb.AddResponse!
	sum := req.A + req.B
	return &pb.AddResponse{Result: sum}, nil
}

func main() {
	// 3. Start the Server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})

	go func() {
		slog.Info("gRPC Server Listing on por 50051")
		s.Serve(lis)
	}()

	// The credentials MUST be passed as the second argument to NewClient!
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	res, err := c.Add(context.Background(), &pb.AddRequest{A: 10, B: 25})
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	slog.Info("Result", slog.Any("Sum", res))
}
