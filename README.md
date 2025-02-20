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

## Setting up JWT Authentication

1. Create a new directory named `auth` and add an `auth.go` file with the following content:
   ```go
   package auth

   import (
       "time"
       "github.com/dgrijalva/jwt-go"
   )

   var jwtKey = []byte("my_secret_key")

   type Claims struct {
       Username string `json:"username"`
       jwt.StandardClaims
   }

   func GenerateJWT(username string) (string, error) {
       expirationTime := time.Now().Add(24 * time.Hour)
       claims := &Claims{
           Username: username,
           StandardClaims: jwt.StandardClaims{
               ExpiresAt: expirationTime.Unix(),
           },
       }

       token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
       tokenString, err := token.SignedString(jwtKey)
       if err != nil {
           return "", err
       }

       return tokenString, nil
   }

   func ValidateJWT(tokenString string) (*Claims, error) {
       claims := &Claims{}

       token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
           return jwtKey, nil
       })

       if err != nil {
           if err == jwt.ErrSignatureInvalid {
               return nil, err
           }
           return nil, err
       }

       if !token.Valid {
           return nil, err
       }

       return claims, nil
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

## Running the gRPC Server with JWT Authentication

1. Run the Go project:
   ```sh
   go run main.go
   ```

2. The gRPC server should start with JWT authentication enabled.

## Login, Registration, Password Reset, Forgot Password, OTP Sending, and OTP Verification

1. Add the following functions to `auth/auth.go`:
   ```go
   func Register(username, password string) error {
       // Hash the password
       hashedPassword, err := HashPassword(password)
       if err != nil {
           return err
       }

       // Store the username and hashed password in your database
       // This is a placeholder, replace with actual database logic
       _ = username
       _ = hashedPassword

       return nil
   }

   func Login(username, password string) (string, error) {
       // Retrieve the hashed password from your database
       // This is a placeholder, replace with actual database logic
       storedHashedPassword := "$2a$14$u1Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q0Q
