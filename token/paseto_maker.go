package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PastoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (*PastoMaker, error) {
	if len(symmetricKey) < chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size, must be at least %d characters", chacha20poly1305.KeySize)
	}

	maker := &PastoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PastoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := CreateNewPayload(username, duration)
	if err != nil {
		return "", err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, err
}

func (maker *PastoMaker) VerifyToken(token string) (Payload, error) {
	payload := Payload{}
	err := maker.paseto.Decrypt(token, maker.symmetricKey, &payload, nil)

	if err != nil {
		return Payload{}, ErrInvalidToken
	}

	err = payload.Valid()

	if err != nil {
		return Payload{}, ErrExpiredToken
	}

	return payload, err
}
