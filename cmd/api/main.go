package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lucasdamasceno96/boxing-list/internal/server/routes"
)

func main() {
	router := gin.Default()

	v1 := router.Group("api/v1")
	routes.SetupRoutes(v1)

	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
