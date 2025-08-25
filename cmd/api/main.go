package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lucasdamasceno96/boxing-list/config"
	"github.com/lucasdamasceno96/boxing-list/internal/server/routes"
)

func main() {
	// Init configs
	err := config.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	router := gin.Default()

	v1 := router.Group("api/v1")
	routes.SetupRoutes(v1)

	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
