package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Boxers CRUD - apenas respondendo "ok" para teste
func GetAllBoxers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func GetBoxer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func CreateBoxer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func UpdateBoxer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func DeleteBoxer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
