package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckPassword(t *testing.T) {
	password := RandomString(6)
	hashed_password, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed_password)
	require.NotEqual(t, password, hashed_password)

	err = CheckPassword(password, hashed_password)
	require.NoError(t, err)

	wrong_password := RandomString(6)
	err = CheckPassword(wrong_password, hashed_password)
	require.Error(t, err)

	hashed_password1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed_password1)
	require.NotEqual(t, password, hashed_password1)

	require.NotEqual(t, hashed_password, hashed_password1)
}
