package main

import (
	"context"
	"log"
	"os"

	pb "GOTests/myrpc/order"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"

	defaultName = "yuhao"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {

		log.Fatal("did not connect: %v", err)

	}

	defer conn.Close()

	c := pb.NewOrderClient(conn)

	name := defaultName

	if len(os.Args) > 1 {

		name = os.Args[0]

	}

	r, err := c.GetOrder(context.Background(), &pb.GetOrderRequest{Id: name})

	if err != nil {

		log.Fatal("could not greet: %v", err)

	}

	log.Printf("订单信息: %s", r.Message)

}
