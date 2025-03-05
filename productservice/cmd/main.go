package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	api "productservice/product"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterProductServiceServer(s, &Server{})
	log.Println("Starting server...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Server struct {
	api.UnimplementedProductServiceServer
}

func (s Server) CheckStock(ctx context.Context, request *api.CheckStockRequest) (*api.CheckStockResponse, error) {
	stock := map[int64]int64{
		101: 2,
		102: 3,
		103: 0,
	}

	var response api.CheckStockResponse

	for _, id := range request.ProductIds {
		response.ItemsState = append(response.ItemsState, &api.ProductState{
			ProductId: id,
			Count:     stock[id],
		})
	}

	return &response, nil
}

func (s Server) mustEmbedUnimplementedProductServiceServer() {
}
