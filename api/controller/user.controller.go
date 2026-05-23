package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kai-zenn/bljr_go_api/api/configs"
	"github.com/kai-zenn/bljr_go_api/api/model"
	"github.com/kai-zenn/bljr_go_api/api/utils"
)

func CreateUser(c *gin.Context) {
  var user model.User
  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  hashedPassword, err := utils.HashPassword(user.Password)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password properly"})
    return
  }

  user.Password = hashedPassword

  if err := configs.DB.Create(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"message": "User created succesfully"})
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

func GetUsers(c *gin.Context) {
	var users []model.User
	if err := configs.DB.Preload("Books").Find(&users).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Username string
		Password string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	var user model.User
	if err := configs.DB.First(&user, userID).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	configs.DB.Model(&user).Updates(model.User{
		Username: body.Username,
		Password: body.Password,
	})

	c.JSON(200, gin.H{
		"user": user,
	})
}
