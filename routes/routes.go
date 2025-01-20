package routes

import (
	"github.com/gin-gonic/gin"
	"oms/handlers"
)

func SetupRoutes(router *gin.Engine) {

	//router := gin.Default()
	api := router.Group("/user")
	api.POST("/register", handlers.RegisterUser)
	api.POST("/login", handlers.Login)
	user := api.Group("/user")

	user.GET("/users", handlers.GetAllUsers)
	user.GET("/user/:id", handlers.GetUserById)
	user.DELETE("/user/:id", handlers.DeleteUser)

	product := api.Group("/products")
	{
		product.POST("/create-product", handlers.CreateProductById)
		product.GET("/get-product/:id", handlers.GetProductByid)
		product.GET("/getallproducts", handlers.GetAllProducts)

	}
}
