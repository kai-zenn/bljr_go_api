package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kai-zenn/bljr_go_api/api/configs"
	"github.com/kai-zenn/bljr_go_api/api/model"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var body struct {
		FirstName string 
		LastName  string 
		Username  string 
		Email     string 
		Password  string 
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(body.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := model.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Username:  body.Username,
		Email:     body.Email,
		Password:  string(hashedPassword),
	}


	if err := configs.DB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	userID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	var user model.User
	if err := configs.DB.Preload("Books").First(&user, userID).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}