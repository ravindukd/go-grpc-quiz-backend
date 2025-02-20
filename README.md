# go-grpc-quiz-backend

Trying out Co Pilot capabilities to implement a full project

## Initializing a Go Project

1. Initialize a new Go module:
   ```sh
   go mod init go-grpc-quiz-backend
   ```

2. Create a `main.go` file in the root directory with the following content:
   ```go
   package main

   import (
       "fmt"
       "net"
       "google.golang.org/grpc"
   )

   func main() {
       fmt.Println("Hello, World!")
   }
   ```

## Setting up Supabase

1. Create a new directory named `supabase` and add a `config.go` file with the following content:
   ```go
   package supabase

   import (
       "github.com/supabase/supabase-go"
   )

   func InitializeSupabase() {
       // Add Supabase configuration settings here
   }
   ```

## Setting up gRPC

1. Create a new directory named `grpc` and add a `server.go` file with the following content:
   ```go
   package grpc

   import (
       "google.golang.org/grpc"
   )

   func StartGRPCServer() {
       // Set up a basic gRPC server here
   }
   ```

## Running the Hello World Function

1. Run the Go project:
   ```sh
   go run main.go
   ```

2. You should see the output:
   ```
   Hello, World!
   ```
