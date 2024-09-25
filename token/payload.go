package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrInvalidSecretKey = errors.New("invalid secret key")
var ErrExpiredToken = errors.New("token is expired")
var ErrInvalidToken = errors.New("token is invalid")

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func CreateNewPayload(username string, duration time.Duration) (Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return Payload{}, err
	}

	payload := Payload{
		ID:        id,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (p Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
