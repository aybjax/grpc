package main

import (
	"context"
	pb "grpc/datafiles/transaction"
	"io"
	"log"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started listening to the server stream!")

	stream, err := client.MakeTransaction(context.Background(), request)

	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v\n", client, err)
	}

	for {
		response, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v\n",client, err)
		}

		log.Printf("Status: %v, Operation: %v\n", response.Status, response.Description)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	
	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewMoneyTransactionClient(conn)

	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	ReceiveStream(client, &pb.TransactionRequest{
		From: from,
		To: to,
		Amount:  amount,
	})
}