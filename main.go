package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", routerHandler)

	router.GET("/hello", helloHandler)

	router.Run(":8081")
}

func routerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "belajar dasar web api dengan lybrary gin",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"nama":          "LM. ANANG BRILYANSYAH",
		"Jenis Kelamin": "Pria",
	})
}
