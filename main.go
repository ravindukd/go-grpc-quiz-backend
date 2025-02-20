package main

import (
	"fmt"
	"net"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, World!")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
