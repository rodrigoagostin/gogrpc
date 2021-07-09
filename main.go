package main

import (
	"log"
	"net"

	"github.com/rodrigoagostin/grpc/grpc/pb"
	"github.com/rodrigoagostin/grpc/grpc/service"
	"github.com/rodrigoagostin/grpc/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var ProductList = model.NewProducts()

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Error to connect %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	productService := service.NewProductGrpcService()
	productService.Products = ProductList
	pb.RegisterProductServiceServer(grpcServer, productService)

	grpcServer.Serve(lis)
}
