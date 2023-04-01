package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var transactions []transaction

func Addtransactions(ctx *gin.Context) {
	if err := json.NewDecoder(ctx.Request.Body).Decode(&transactions); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	if isInFuture(transactions[0]) {
		ctx.JSON(http.StatusUnprocessableEntity, nil)
	} else if isExpired(transactions) {
		ctx.JSON(http.StatusNoContent, nil)
	}
	// transactions = append(transactions, transaction)
	ctx.JSON(http.StatusCreated, transactions)
}

func GetStatistics(ctx *gin.Context) {

}

func Deletetransactions(ctx *gin.Context) {
	transactions = []transaction{}
	ctx.JSON(http.StatusNoContent, nil)
}

func AddLocation(ctx *gin.Context) {

}
