package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kai-zenn/bljr_go_api/api/configs"
	"github.com/kai-zenn/bljr_go_api/api/model"
)

func CreateBook(c *gin.Context) {
	// Get data from req body
	var body struct {
		Title string
		Year int
		UserId uuid.UUID 
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	book := model.Book{Title: body.Title, UserID: body.UserId, Year: body.Year}

	if err := configs.DB.Create(&book).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// fetch the created book with associated user
	configs.DB.Preload("User").First(&book, "id = ?", book.ID)
	
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