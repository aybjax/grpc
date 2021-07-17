package main

import (
	"context"
	pb "grpc/datafiles/transaction"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {}

func (s *server) MakeTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	log.Printf("Got request for money Transfer...\n")
	log.Printf("Amount: %f, From A/C: %s, To A/C: %s", in.Amount, in.From, in.To)

	return &pb.TransactionResponse{Confirmation: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}