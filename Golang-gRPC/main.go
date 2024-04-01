package main

import (
	"context"
	"log"
	"net"

	"gRPC-apps/invoicer"

	"google.golang.org/grpc"
)

type MyInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s *MyInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	// Process the request and return the response
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte(req.To),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50051") // Address format is "host:port"
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		log.Printf("Server is serving")
	}

	server := grpc.NewServer()
	service := &MyInvoicerServer{}

	invoicer.RegisterInvoicerServer(server, service)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
