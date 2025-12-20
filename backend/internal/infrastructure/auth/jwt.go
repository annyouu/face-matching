package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTServiceInterface interface {
	// トークンを生成する
	GenerateToken(userID string) (string, error)
	// トークンを検証する
	ValidateToken(tokenString string) (string, error)
}

type jwtService struct {
	secretKey []byte
}

func NewJWTService(secret string) JWTServiceInterface {
	return &jwtService{
		secretKey: []byte(secret),
	}
}

type customClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *jwtService) GenerateToken(userID string) (string, error) {
	claims := &customClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}
	
	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}
	return claims.UserID, nil
}