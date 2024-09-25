package api

import (
	"fmt"
	"net/http"

	db "github.com/SohamKanji/simple-bank-project/db/sqlc"
	"github.com/SohamKanji/simple-bank-project/token"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromAccount, valid := server.isCurrencyValidForAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("currency mismatch")))
		return
	}

	_, valid = server.isCurrencyValidForAccount(ctx, req.ToAccountID, req.Currency)

	if !valid {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("currency mismatch")))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(token.Payload)

	if fromAccount.Owner != authPayload.Username {
		ctx.JSON(http.StatusForbidden, errorResponse(fmt.Errorf("from account does not belong to the authenticated user")))
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	account, err := server.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
