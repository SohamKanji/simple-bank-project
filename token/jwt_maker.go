package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const min_secret_key_size = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (*JWTMaker, error) {
	if len(secretKey) < min_secret_key_size {
		return nil, ErrInvalidSecretKey
	}

	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a new JWT token
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := CreateNewPayload(username, duration)

	if err != nil {
		return "", &Payload{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload.Username,
		"exp": payload.ExpiredAt.Unix(),
		"iat": payload.IssuedAt.Unix(),
		"id":  payload.ID.String(),
	})

	signed_token, err := token.SignedString([]byte(maker.secretKey))

	return signed_token, &payload, err
}

// VerifyToken checks if the token is valid or not

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	jwtToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return &Payload{}, ErrExpiredToken
		}
		return &Payload{}, ErrInvalidToken
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok || !jwtToken.Valid {
		return &Payload{}, ErrInvalidToken
	}

	if exp, ok := claims["exp"].(float64); ok {
		expirationTime := time.Unix(int64(exp), 0)

		if time.Now().After(expirationTime) {
			return &Payload{}, ErrExpiredToken
		}
	}

	return &Payload{
		ID:        uuid.MustParse(claims["id"].(string)),
		Username:  claims["sub"].(string),
		IssuedAt:  time.Unix(int64(claims["iat"].(float64)), 0),
		ExpiredAt: time.Unix(int64(claims["exp"].(float64)), 0),
	}, nil
}
