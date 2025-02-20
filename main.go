package main

import (
	"fmt"
	"net"
	"google.golang.org/grpc"
	"grpc"
	"auth"
	"supabase"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
		return
	}

	supabaseClient := supabase.InitializeSupabase()
	auth.SetSupabaseClient(supabaseClient)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc.jwtAuthInterceptor))
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
