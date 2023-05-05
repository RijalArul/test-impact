package main

import (
	"test-impact/server/databases"

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
func main() {
	g := gin.Default()

	databases.StartDB()

	g.Use(CorsMiddleware())

	g.Run()
}
