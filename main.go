package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", routerHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler) //this is path variable
	router.GET("/query", queryHandler)            //this is query string

	router.Run()
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

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}
