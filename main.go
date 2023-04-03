package main

import (
	"banking/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/transactions", controllers.Addtransactions)

	router.POST("/location", controllers.AddLocation)

	router.GET("/statistics", controllers.GetStatistics)

	router.DELETE("/transactions", controllers.Deletetransactions)

	router.PUT("/resetLocation", controllers.ResetLocation)

	router.Run("localhost:8000")
}
