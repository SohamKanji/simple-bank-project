package token

import (
	"testing"
	"time"

	"github.com/SohamKanji/simple-bank-project/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	secretKey := util.RandomString(32)

	maker, err := NewPasetoMaker(secretKey)
	require.NoError(t, err)

	username := util.RandomOwner()

	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	secretKey := util.RandomString(32)

	maker, err := NewPasetoMaker(secretKey)
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := -time.Minute

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Equal(t, Payload{}, payload)

}

func TestInvalidPasetoToken(t *testing.T) {
	secretKey := util.RandomString(32)

	maker, err := NewPasetoMaker(secretKey)
	require.NoError(t, err)

	username := util.RandomOwner()

	duration := time.Minute

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	invalid_token := token + "invalid"

	payload, err := maker.VerifyToken(invalid_token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Equal(t, Payload{}, payload)
}
