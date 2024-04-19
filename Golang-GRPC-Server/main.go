package main

import (
	"context"
	"fmt"
	"log"
	"net"
	// "github.com/gocql/gocql"
	"google.golang.org/grpc"
	pb "grpc-server-test/protobuf"
	cass "grpc-server-test/connection"
	qs "grpc-server-test/cassandrahelper"

)

const (
	// Port for gRPC server to listen to
	PORT = ":50051"
)

type server struct {
	pb.UnimplementedSmrSyncServer
}

func (s *server) SyncData(ctx context.Context, in *pb.DataRequest) (*pb.DataResponse, error) {
	// Log the received message
	log.Printf("Received SyncData request with data: %v", in)

	// Update data in Cassandra
	if err := qs.CassandraQueryNumber(ctx, in); err != nil {
		return nil, fmt.Errorf("failed to update data in Cassandra: %v", err)
	}

	// Return success response
	return &pb.DataResponse{Success: true}, nil
}

func (s *server) SyncNumber(ctx context.Context, in *pb.DataNumber) (*pb.DataResponse, error) {
	// Log the received message
	// log.Printf("Received SyncNumber request with data: %v", in)

	// Your implementation for the SyncNumber RPC method
	// Example: Perform synchronization logic here
	return &pb.DataResponse{Success: true}, nil
}

func main() {
	// Establish Cassandra connection
	cass.ConnectionCassandra()
	defer cass.CloseCassandra()

	// Set up gRPC server
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*50), // Increase the maximum receive message size to 50 MB
	)
	pb.RegisterSmrSyncServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}