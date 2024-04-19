package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	mh "sync-databases-cass-mysql/helper"
	pb "sync-databases-cass-mysql/protobuf"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const (
	PORT = ":50051"
	// BATCH_SIZE = 10000 //Batching per 10000
)

type server struct {
	pb.UnimplementedSmrSyncServer
}

func main() {
	// KOneksi ke GRPC
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("HOST")
	portStr := os.Getenv("PORT")

	// Convert port string to int
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Error converting port to integer: %v", err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Client
	client := pb.NewSmrSyncClient(conn)

	// Context and Function untuk GRPC
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Goroutin
	ticker := time.NewTicker(1 * time.Minute) // Adjust the interval as needed
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done(): // Channel context when done
			return
		case <-ticker.C: // Interval untuk channel
			// Synchronize data
			err := mh.SyncDataAndAvoidParallel(ctx, client)
			if err != nil {
				log.Printf("Error synchronizing data: %v", err)
			}
		}
	}
}