package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	transact     transaction
	Transactions []transaction
	locate       location
)

func Addtransactions(ctx *gin.Context) {
	if locate.City != "" {
		if err := json.NewDecoder(ctx.Request.Body).Decode(&transact); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			log.Fatal(err)
			return
		}
		if isInFuture(transact) {
			ctx.JSON(http.StatusUnprocessableEntity, nil)
		} else if isExpired(transact) {
			ctx.JSON(http.StatusNoContent, nil)
		} else {
			Transactions = append(Transactions, transact)
			ctx.JSON(http.StatusCreated, Transactions)
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, nil)
	}

}

// GetStatistics will provide statistics
func GetStatistics(ctx *gin.Context) {
	if locate.City != "" {
		translastsix := []transaction{}
		for _, v := range Transactions {
			if !isExpired(v) {
				translastsix = append(translastsix, v)
			}
		}
		if !isTrasact(translastsix) {
			ctx.JSON(http.StatusOK, "No transactions from last 60 seconds.....")
		} else {
			statistic := statistic(translastsix)
			ctx.JSON(http.StatusOK, statistic)
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, nil)
	}

}

func Deletetransactions(ctx *gin.Context) {
	if locate.City != "" {
		Transactions = []transaction{}
		ctx.JSON(http.StatusNoContent, nil)
	} else {
		ctx.JSON(http.StatusUnauthorized, nil)
	}
}

func AddLocation(ctx *gin.Context) {
	if err := json.NewDecoder(ctx.Request.Body).Decode(&locate); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
		return
	}
}

func ResetLocation(ctx *gin.Context) {
	locate = location{}
	ctx.JSON(http.StatusNoContent, nil)
}
