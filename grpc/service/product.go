package service

import (
	"context"

	"github.com/rodrigoagostin/grpc/grpc/pb"
	"github.com/rodrigoagostin/grpc/model"
)

type ProductGrpcService struct {
	pb.UnimplementedProductServiceServer
	Products *model.Products
}

func NewProductGrpcService() *ProductGrpcService {
	return &ProductGrpcService{}
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.Product) (*pb.ProductResult, error) {
	product := model.NewProduct()
	product.Name = in.GetName()
	p.Products.Add(product)

	return &pb.ProductResult{
		Id:   product.ID,
		Name: product.Name,
	}, nil
}
