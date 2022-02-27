package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// JWTMaker is JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker create a new JWTMaker
func NewJWTMaker(secretKey string) Maker {
	return &JWTMaker{secretKey}
}

// CreateToken create jwt token string
func (maket *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, _ := NewPayload(username, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maket.secretKey))
}

// Verify check if token is valid jwt token or not
func (maket *JWTMaker) Verify(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(maket.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		if _, ok := err.(*jwt.ValidationError); !ok {
			return nil, errors.New("token expired")
		}
		return nil, errors.New("invalid token")
	}
	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, errors.New("invalid token")
	}
	return payload, nil
}
