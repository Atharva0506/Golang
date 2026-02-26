# The Senior Engineer's Guide to gRPC in Go

## 1. What is gRPC and Why Do We Need It?
In a standard Microservice architecture, **Service A** talks to **Service B** using an HTTP REST API and sends data using **JSON**.
This creates a massive bottleneck for companies like Google, Netflix, and Uber:
1. **JSON is huge:** The word `"false"` takes 5 bytes. The number `1000` takes 4 bytes.
2. **JSON is slow:** CPU power is wasted parsing the JSON string into native code structures.
3. **No Safety:** If Service B changes an API route or a struct field, Service A doesn't know until it crashes in production.

**gRPC solves this using Protobuf (Protocol Buffers):**
Instead of sending bulky text strings, gRPC compresses your data into **pure binary** (`010101`) before sending it over TCP. It is 7x to 10x faster than JSON APIs, uses incredibly low bandwidth, and guarantees that Server and Client code is perfectly synced.

---

## 2. Installation on Windows
To use gRPC in Go, you need three things:
1. The **protoc** compiler (A C++ program built by Google that reads [.proto](file:///d:/Practice/Go/03_Advanced/14_grpc/proto/calculator.proto) files).
2. The **protoc-gen-go** plugin (Tells the compiler how to generate Go structs).
3. The **protoc-gen-go-grpc** plugin (Tells the compiler how to generate the server/client endpoints).

### Step 1: Install the Compiler
1. Go to the [Protobuf Releases Page on GitHub](https://github.com/protocolbuffers/protobuf/releases).
2. Download the `protoc-XX.X-win64.zip` file.
3. Extract it, and put the `protoc.exe` file inside your `$GOPATH/bin` folder (or somewhere in your Windows Environment `$PATH`).

### Step 2: Install the Go Plugins
Open your terminal and run:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---

## 3. Creating the [.proto](file:///d:/Practice/Go/03_Advanced/14_grpc/proto/calculator.proto) File (The Rulebook)
In gRPC, you write a [.proto](file:///d:/Practice/Go/03_Advanced/14_grpc/proto/calculator.proto) file. This acts as the absolute source of truth for your API.

Create [proto/calculator.proto](file:///d:/Practice/Go/03_Advanced/14_grpc/proto/calculator.proto):
```protobuf
syntax = "proto3";

// Defines the Go package path
option go_package = "grpc_demo/proto";

// The Request Message
message AddRequest {
    int32 a = 1;
    int32 b = 2;
}

// The Response Message
message AddResponse {
    int32 result = 1;
}

// The API Endpoints (Procedures)
service CalculatorService {
    rpc Add(AddRequest) returns (AddResponse);
}
```

---

## 4. Compiling the Code
Once your [.proto](file:///d:/Practice/Go/03_Advanced/14_grpc/proto/calculator.proto) file is written, you open your terminal and tell the compiler to build your API for you!

Run this command from the root of your project:
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/calculator.proto
```
This instantly creates two files: `calculator.pb.go` and `calculator_grpc.pb.go`. **Never edit these files!** They contain thousands of lines of bulletproof networking code.

---

## 5. Writing the Go Code
Now, you just initialize your module, write the business logic, and test it!

### Step 1: Initialize Go
```bash
go mod init my_app
go mod tidy
```

### Step 2: The Logic ([main.go](file:///d:/Practice/Go/01_Beginner_Test/main.go))
```go
package main

import (
	"context"
	"log"
	"net"

	pb "my_app/proto" // Import the generated code!

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 1. We embed the generated UnimplementedServer to ensure we satisfy the interface
type server struct {
	pb.UnimplementedCalculatorServiceServer
}

// 2. We write the actual business logic for the Add method we defined in the .proto file
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B
	return &pb.AddResponse{Result: sum}, nil
}

func main() {
	// --- START THE SERVER ---
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})

    // We run it in a goroutine so it doesn't block the rest of our `main` function
	go func() {
		log.Println("gRPC Server Listening on port 50051")
		s.Serve(lis)
	}()

	// --- BUILD THE CLIENT AND TEST IT ---
    
    // Connect to the server. (We use insecure credentials because we don't have SSL/TLS on localhost)
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

    // Instantiate the native gRPC client
	c := pb.NewCalculatorServiceClient(conn)

    // Call the server! It looks like a local function, but it's sending binary over TCP!
	res, err := c.Add(context.Background(), &pb.AddRequest{A: 10, B: 25})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server responded with: %v", res.Result)
}
```

## 6. The 4 Types of gRPC Requests

gRPC supports 4 different types of communication. We just built the first one!

1. **Unary (Simple RPC):** Client sends 1 request, Server sends 1 response. *(What we just built)*
2. **Server Streaming:** Client sends 1 request, Server returns a continuous stream of responses. *(Example: Live stock market ticker feed)*
3. **Client Streaming:** Client sends a continuous stream of data, Server returns 1 single response when finished. *(Example: Uploading a massive 4GB video file chunk-by-chunk)*
4. **Bidirectional Streaming:** Client and Server both send continuous streams of data back and forth simultaneously. *(Example: Real-time multiplayer video game)*
