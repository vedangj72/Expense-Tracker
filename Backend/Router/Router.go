package router

import (
	controller "ExpenseTacker/Controller"
	middleware "ExpenseTacker/Middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router and defines all the routes.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/register", controller.RegisterUserController)
	router.POST("/login", controller.LoginUserController)
	router.GET("/register", controller.GetAllRegisterUser)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running for expense tracker",
		})
	})

	// Group for authorized routes
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/total", controller.GetTotalAmountController)
		authorized.POST("/total", controller.PostTotalAmountController)
		authorized.POST("/totaltransaction", controller.PostTotalAmountController)
		authorized.GET("/totaltransaction", controller.GetTransactionsController)
	}

	return router
}
