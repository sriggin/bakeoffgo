package main

import (
	ts "bakeoffgo/trafficscanner"
	"context"
	"fmt"
	"github.com/bourbaki/catboost-go"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = 50051
)

type trafficScannerServer struct {
	ts.UnimplementedTrafficScannerServer
	model *catboost.Model
}

func (s *trafficScannerServer) IsTrafficValid(ctx context.Context, in *ts.TrafficScanRequest) (*ts.TrafficScanResponse, error) {
	response := ts.TrafficScanResponse{
		IsValid: true,
	}

	return &response, nil
}

func main() {
	model, err := catboost.LoadFullModelFromFile("model.cbm")
	if err != nil {
		log.Fatalf("Failed to load model: %v", err)
	}

	server := trafficScannerServer{model: model}

	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	ts.RegisterTrafficScannerServer(grpcServer, &server)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve TrafficScanner service on listener: %v", err)
	}
}
