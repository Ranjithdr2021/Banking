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
	res          reset
)

func Addtransactions(ctx *gin.Context) {
	if res.Reset == true || locate.City != "" {
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
			ctx.JSON(http.StatusCreated, "SUCCESS")
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, nil)
	}

}

// GetStatistics will provide statistics
func GetStatistics(ctx *gin.Context) {
	if res.Reset == true || locate.City != "" {
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
	if res.Reset == true || locate.City != "" {
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
	ctx.JSON(http.StatusOK, "Location added Successfully...")
}

func ResetLocation(ctx *gin.Context) {
	locate = location{}
	res.Reset = true
	ctx.JSON(http.StatusOK, "location reset Successfull...")
}
