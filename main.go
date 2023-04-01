package main

import (
	"banking/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/transactions", controllers.Addtransactions)

	// router.PUT("/update-status/:O_id", controllers.UpdateStatus)

	router.GET("/statistics", controllers.GetStatistics)

	// router.GET("/GetOrdersById/:O_id", controllers.GetOrdersById)

	// router.GET("/orders", controllers.Orders)

	router.Run("localhost:8000")
}
