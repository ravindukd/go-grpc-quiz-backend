package auth

import (
	"testing"
	"time"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func TestGenerateJWT(t *testing.T) {
	username := "testuser"
	token, err := GenerateJWT(username)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	claims := &Claims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if claims.Username != username {
		t.Fatalf("Expected username %v, got %v", username, claims.Username)
	}
}

func TestValidateJWT(t *testing.T) {
	username := "testuser"
	token, err := GenerateJWT(username)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	claims, err := ValidateJWT(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if claims.Username != username {
		t.Fatalf("Expected username %v, got %v", username, claims.Username)
	}
}

func TestHashPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "password123"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !CheckPasswordHash(password, hashedPassword) {
		t.Fatalf("Expected password to match hash")
	}
}

func TestRegister(t *testing.T) {
	username := "testuser"
	password := "password123"

	err := Register(username, password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var result []map[string]interface{}
	err = supabaseClient.From("users").Select("password").Eq("username", username).Execute(&result)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) == 0 {
		t.Fatalf("Expected user to be registered")
	}
}

func TestLogin(t *testing.T) {
	username := "testuser"
	password := "password123"

	err := Register(username, password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	token, err := Login(username, password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	claims, err := ValidateJWT(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if claims.Username != username {
		t.Fatalf("Expected username %v, got %v", username, claims.Username)
	}
}
