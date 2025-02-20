package grpc

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"time"
	"github.com/dgrijalva/jwt-go"
	"auth"
	"supabase"
)

type mockServer struct{}

func (s *mockServer) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello " + in.Name}, nil
}

func TestSayHello(t *testing.T) {
	s := &mockServer{}
	req := &HelloRequest{Name: "World"}
	resp, err := s.SayHello(context.Background(), req)
	if err != nil {
		t.Fatalf("SayHello() error = %v", err)
	}
	if resp.Message != "Hello World" {
		t.Fatalf("SayHello() = %v, want %v", resp.Message, "Hello World")
	}
}

func TestJwtAuthInterceptor(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "valid-token"))
	req := &HelloRequest{}
	info := &grpc.UnaryServerInfo{}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return &HelloReply{Message: "Hello"}, nil
	}

	_, err := jwtAuthInterceptor(ctx, req, info, handler)
	if err != nil {
		t.Fatalf("jwtAuthInterceptor() error = %v", err)
	}
}
