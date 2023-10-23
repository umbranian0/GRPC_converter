package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/umbranian0/GRPC_converter/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr     = flag.String("addr", "localhost:50051", "the address to connect to")
	operand1 = flag.Int("op1", 2, "1st operand")
	operand2 = flag.Int("op2", 2, "2nd operand")

	operand1int32 = int32(*operand1)
	operand2int32 = int32(*operand2)
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &pb.Input{Operand1: operand1int32, Operand2: operand2int32})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Add result: %v", r.GetResult())
}
