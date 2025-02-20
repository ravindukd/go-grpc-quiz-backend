package main

import (
	"testing"
	"net"
	"google.golang.org/grpc"
	"grpc"
	"auth"
	"supabase"
	"os"
	"github.com/joho/godotenv"
)

func TestMainFunction(t *testing.T) {
	go main()

	// Wait for the server to start
	time.Sleep(2 * time.Second)

	// Check if the server is running
	conn, err := net.Dial("tcp", ":50051")
	if err != nil {
		t.Fatalf("failed to connect to server: %v", err)
	}
	conn.Close()
}
