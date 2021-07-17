package main

import (
	"fmt"
	pb "grpc/datafiles/transaction"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
	noOfSteps = 3
)

type server struct {}

func (s *server) MakeTransaction(in *pb.TransactionRequest,
	stream pb.MoneyTransaction_MakeTransactionServer) error {
		log.Println("Got request for money transfer")
		log.Printf("Amount: $%f, From A/C: %s, To A/C: %s\n",
		in.Amount, in.From, in.To)
	
	// send stream
	for i:=0; i < noOfSteps; i++ {
		time.Sleep(time.Second * 2)

		if err := stream.Send(&pb.TransactionResponse{Status: "Good",
		Step: int32(i),
		Description: fmt.Sprintf("Description of step %d\n", int32(i),)});
		err != nil {
			log.Fatalf("%v.Send(%v) = %v\n", stream, "status", err)
		}
	}

	log.Printf("Successfully transfered amount $%v frm %v to %v\n",
						in.Amount, in.From, in.To)

	return nil;
}


func main() {
	lis, err := net.Listen("tcp", port)
	
	if err != nil {
		log.Fatalf("Failed to listen to %v\n", err)
	}

	s := grpc.NewServer()

	pb.RegisterMoneyTransactionServer(s, &server{})

	reflection.Register(s)
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}