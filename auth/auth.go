package auth

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/smtp"
	"errors"
	"github.com/supabase/supabase-go"
	"os"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))
var otpMap = make(map[string]string)
var supabaseClient = supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(username, password string) error {
	// Hash the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	// Store the username and hashed password in Supabase
	_, err = supabaseClient.From("users").Insert(map[string]interface{}{
		"username": username,
		"password": hashedPassword,
	}).Execute()
	if err != nil {
		return err
	}

	return nil
}

func Login(username, password string) (string, error) {
	// Retrieve the hashed password from Supabase
	var result []map[string]interface{}
	err := supabaseClient.From("users").Select("password").Eq("username", username).Execute(&result)
	if err != nil {
		return "", err
	}

	if len(result) == 0 {
		return "", errors.New("user not found")
	}

	storedHashedPassword := result[0]["password"].(string)

	// Check if the provided password matches the stored hashed password
	if !CheckPasswordHash(password, storedHashedPassword) {
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token, err := GenerateJWT(username)
	if err != nil {
		return "", err
	}

	return token, nil
}
