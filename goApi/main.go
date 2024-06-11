package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Romeo and julet", Author: "Shakeshpear", Quantity: 4},
	{ID: "2", Title: "Tempest", Author: "Shakeshpear", Quantity: 8},
	{ID: "3", Title: "Metamorphosois", Author: "Franz Kafka", Quantity: 2},
	{ID: "4", Title: "Sisphus", Author: "Albert Campus", Quantity: 1},
	{ID: "5", Title: "Castle", Author: "Franz Kafka", Quantity: 7},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBooks(c *gin.Context) {
	var newbook book

	if err := c.BindJSON(&newbook); err != nil {
		return
	}
	books = append(books, newbook)
	c.IndentedJSON(http.StatusCreated, newbook)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
	}

	book, err := getBookID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}
	if book.Quantity < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No book avaliable "})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
	}

	book, err := getBookID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBooks)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8888")
}
