package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdamasceno96/boxing-list/internal/server/handlers"
)

func SetupRoutes(r *gin.RouterGroup) {
	r.GET("health", handlers.Health)

	boxerRoutes := r.Group("/boxers")
	// Adiciona rotas diretamente no grupo
	boxerRoutes.GET("/all", handlers.GetAllBoxers)   // List all
	boxerRoutes.GET("", handlers.GetBoxer)           // List by name or Id
	boxerRoutes.POST("", handlers.CreateBoxer)       // Create
	boxerRoutes.PUT("/:id", handlers.UpdateBoxer)    // Update
	boxerRoutes.DELETE("/:id", handlers.DeleteBoxer) // Delete
}
