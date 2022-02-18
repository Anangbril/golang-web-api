package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"subject": "clue olahraga",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"nama":          "LM. ANANG BRILYANSYAH",
			"Jenis Kelamin": "Pria",
		})
	})

	router.Run(":8081")
}
