package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Addtransactions(ctx *gin.Context) {
	transaction := transaction{}
	json.NewDecoder(ctx.Request.Body).Decode(&transaction)
	if isExpired(transaction) {
		ctx.JSON(http.StatusUnprocessableEntity, nil)
	}
	ctx.JSON(http.StatusCreated, transaction)
}
