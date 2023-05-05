package routes

import (
	"test-impact/server/databases"
	"test-impact/server/handlers"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cors.New(cors.Options{
			AllowedOrigins: []string{"http://localhost:3000", "https://www.google.com"},
			AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Content-Type"},
			Debug:          true,
		})
	}
}
func Routes() *gin.Engine {
	r := gin.Default()

	r.Use(CorsMiddleware())
	getDB := databases.GetDB()

	handlerProducts := handlers.NewProductHandler(getDB)
	routerProducts := r.Group("/products")

	{
		routerProducts.GET("/", handlerProducts.Products)
		routerProducts.POST("/create", handlerProducts.Create)
		routerProducts.GET("/:product_id", handlerProducts.Product)
		routerProducts.PUT("/update/:product_id", handlerProducts.Update)
		routerProducts.DELETE("/delete/:product_id", handlerProducts.Delete)
	}

	r.Run()
	return r

}
