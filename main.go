package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", routerHandler)
	v1.GET("/hello", helloHandler)
	v1.GET("/books/:id/:title", booksHandler) //this is path variable
	v1.GET("/query", queryHandler)            //this is query string

	v1.POST("/books", postBooksHandler)

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

type BookInput struct {
	Title string      `json:"Title" binding:"required"`
	Price json.Number `json:"Price" binding:"required,number"`
	// SubTitle string      `json:"sub_title"`
}

func postBooksHandler(c *gin.Context) {
	var bookInput BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on fields %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
		// "SubTitle": bookInput.SubTitle,
	})
}
