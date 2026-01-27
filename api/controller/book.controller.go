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
		Title string `json:"title"`
		Year int `json:"year"`
		UserId uuid.UUID `json:"user_id"`
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
	// mapping ke field author dengan username dari user
	book.Author = book.User.Username

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
	configs.DB.Preload("User").Find(&book, id)
	
	// return books
	c.JSON(200, gin.H{
		"book": book,
	})
}

func GetAllBooks(c *gin.Context) {
	var books []model.Book
	if err := configs.DB.Preload("User").Find(&books).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// map author field
	for i := range books {
		books[i].Author = books[i].User.Username
	}
	
	c.JSON(200, gin.H{
		"books": books,
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get a single book that we want to update
	var book model.Book
	if err := configs.DB.First(&book, "id =?", id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "error book not found",
		})
		return
	}

	// update the book fields
	configs.DB.Model(&book).Updates(model.Book{
		Title: body.Title,
	})

	// preload user data
	configs.DB.Preload("User").First(&book, "id = ?", book.ID)
	book.Author = book.User.Username

	// return response
	// c.JSON(200, gin.H{
	// 	"book": book,
	// })
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	configs.DB.Delete(&model.Book{}, id)

	c.JSON(200, gin.H{
		"message": "Book deleted successfully",
	})
}