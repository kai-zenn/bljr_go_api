package controller

import (
	"bljr_go_api/api/configs"
	"bljr_go_api/api/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	// Get data from req body
	var body struct {
		Title string
		Author string
		Year int
	}
	c.Bind(&body)
	// Create a book

	book := model.Book{Title: body.Title, Author: body.Author, Year: body.Year}
	result := configs.DB.Create(&book)

	if result.Error != nil {
		c.Status(400)
		fmt.Println(result.Error)
		return
	}
	// Return it

	c.JSON(200, gin.H{
	  "book": book,
	})
}

func GetBook(c *gin.Context) {
	// get id from url param
	id := c.Param("id")

	// get books from database
	var book []model.Book
	configs.DB.Find(&book, id)

	// return books
	c.JSON(200, gin.H{
		"book": book,
	})
}

func GetAllBooks(c *gin.Context) {
	var books []model.Book
	configs.DB.Find(&books)
	c.JSON(200, gin.H{
		"books": books,
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Author string
		Year int
	}
	c.Bind(&body)

	// Get a single book that we want to update
	var book model.Book
	configs.DB.Find(&book, id)

	// update the book fields
	configs.DB.Model(&book).Updates(model.Book{
		Title: body.Title,
		Author: body.Author,
		Year: body.Year,
	})

	// return response
	c.JSON(200, gin.H{
		"book": book,
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	configs.DB.Delete(&model.Book{}, id)

	c.JSON(200, gin.H{
		"message": "Book deleted successfully",
	})
}