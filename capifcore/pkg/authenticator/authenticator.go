package authenticator

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func generateSecretKey(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

var jwtKey []byte

func init() {
	var err error
	key, err := generateSecretKey(32)
	if err != nil {
		log.Fatalf("Failed to generate JWT key: %v", err)
	}
	jwtKey = []byte(key)
}

func GenerateUserCredential(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateUserCredential(tokenString string) (*jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func HelloWorld() {
	log.Print("Hello, World")
}

const simpleUserToken = "simpleUserToken"
const simpleRappToken = "simpleRappToken"

// User token
// GenerateToken 返回一个简单的静态令牌
func GenerateToken() string {
	return simpleUserToken
}

// ValidateToken 检查提供的令牌是否与我们的简单令牌匹配
func ValidateToken(tokenString string) (bool, error) {
	if tokenString == simpleUserToken {
		return true, nil
	}
	return false, errors.New("invalid token")
}

// rApp token
func GenerateRAppToken() string {
	return simpleRappToken
}

// ValidateToken 检查提供的令牌是否与我们的简单令牌匹配
func ValidateRAppToken(tokenString string) (bool, error) {
	if tokenString == simpleRappToken {
		return true, nil
	}
	return false, errors.New("invalid token")
}
