package main

import (
	"context"
	"log"
	"net"

	pb "github.com/umbranian0/GRPC_converter/calculator/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func main() {

	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) Add(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.Nums)
	var tot int32
	for _, v := range in.Nums {
		tot += v
	}
	return &pb.Reply{Num: tot}, nil
}

func (s *server) Sub(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.Nums)
	var tot int32
	for _, v := range in.Nums {
		tot -= v
	}
	return &pb.Reply{Num: tot}, nil
}

func (s *server) Mul(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.Nums)
	var tot int32
	tot = 1
	for _, v := range in.Nums {

		tot = (tot + v) * v
	}
	return &pb.Reply{Num: tot}, nil
}

func (s *server) Div(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.Nums)
	var tot int32

	for _, v := range in.Nums {

		tot = (tot + v) / v
	}
	return &pb.Reply{Num: tot}, nil
}
