package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "GOTests/myrpc/order"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {

	return &pb.GetOrderResponse{Message: "订单信息是" + in.Id}, nil

}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderServer(s, &server{})
	s.Serve(lis)
}
