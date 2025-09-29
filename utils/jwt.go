package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTUtil helper for JWT tokens
type JWTUtil struct {
	Secret string
	ExpMin int
}

// NewJWTUtil create new JWT util
func NewJWTUtil(secret string, expMin int) *JWTUtil {
	return &JWTUtil{
		Secret: secret,
		ExpMin: expMin,
	}
}

// GenerateToken creates token with subject=userID
func (j *JWTUtil) GenerateToken(subject string, duration time.Duration) (string, error) {
	if duration == 0 {
		duration = time.Duration(j.ExpMin) * time.Minute
	}
	claims := &jwt.RegisteredClaims{
		Subject:   subject,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(j.Secret))
}

// ParseToken verifies JWT and returns claims
func (j *JWTUtil) ParseToken(tokenStr string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
