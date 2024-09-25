package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/SohamKanji/simple-bank-project/token"
	"github.com/SohamKanji/simple-bank-project/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func addAuthorization(t *testing.T, request *http.Request, maker token.Maker, username string, duration time.Duration, authorizationHeaderKey string) {
	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)

	request.Header.Set(authorizationHeaderKey, fmt.Sprintf("Bearer %s", token))
}

func TestMiddleware(t *testing.T) {
	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, maker token.Maker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, maker token.Maker) {
				addAuthorization(t, request, maker, util.RandomOwner(), time.Minute, authorizationHeaderKey)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, maker token.Maker) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidAuthorizationHeaderKey",
			setupAuth: func(t *testing.T, request *http.Request, maker token.Maker) {
				addAuthorization(t, request, maker, util.RandomOwner(), time.Minute, "invalid")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidToken",
			setupAuth: func(t *testing.T, request *http.Request, maker token.Maker) {
				request.Header.Set(authorizationHeaderKey, fmt.Sprintf("Bearer %s", "invalid-token"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "ExpiredToken",
			setupAuth: func(t *testing.T, request *http.Request, maker token.Maker) {
				addAuthorization(t, request, maker, util.RandomOwner(), -time.Minute, authorizationHeaderKey)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			server := NewTestServer(t, nil)
			server.router.GET("/auth", AuthMiddleware(server.maker), func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{})
			})
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "/auth", nil)

			tc.setupAuth(t, request, server.maker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
