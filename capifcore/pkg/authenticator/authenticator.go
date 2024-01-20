package authenticator

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

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

// rApp token
func GenerateSubToken() string {
	return simpleRappToken
}

// ValidateToken 检查提供的令牌是否与我们的简单令牌匹配
func ValidateSubToken(tokenString string) bool {
	if tokenString == simpleRappToken {
		return true
	}
	return false
}

func getJWTKey(apiId string) string {
	return apiId
}

func GenerateSubscribeToken(apiId string) (string, error) {
	jwtKey := getJWTKey(apiId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"apiId": apiId,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateSubscribeToken(tokenString, apiId string) (*jwt.Token, error) {
	jwtKey := getJWTKey(apiId)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}
