package router

import (
	controller "ExpenseTacker/Controller"
	middleware "ExpenseTacker/Middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router and defines all the routes.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// CORS middleware configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Allow all origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	// Public routes (no authentication required)
	router.POST("/register", controller.RegisterUserController)
	router.POST("/login", controller.LoginUserController)
	router.GET("/register", controller.GetAllRegisterUser)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running for expense tracker",
		})
	})

	// Authorized routes (protected by middleware)
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware()) // Middleware to check authentication
	{
		authorized.GET("/total", controller.GetTotalAmountController)
		authorized.POST("/total", controller.PostTotalAmountController)
		authorized.GET("/transactions", controller.GetTransactionsController)
		authorized.POST("/transactions", controller.PostTransactionController)
	}

	return router
}
