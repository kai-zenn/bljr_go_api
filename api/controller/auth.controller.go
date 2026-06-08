package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"github.com/kai-zenn/bljr_go_api/api/configs"
	"github.com/kai-zenn/bljr_go_api/api/model"
	"github.com/kai-zenn/bljr_go_api/api/utils"
)

type LoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginUserHandler(c *gin.Context) {
  var input LoginDTO 
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  
  var user model.User
  if err := configs.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{
      "error": "Invalid Email or Password",
    })
    return
  }

  if !utils.CheckPasswordHash(input.Password, user.Password) {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
    return
  }
  
  token, err := utils.CreateToken(user.ID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": "Failed to generate token",
    })
    return
  }
  
  c.JSON(http.StatusOK, gin.H{
    "token": token,
  })
}

func RegisterUserHandler(c *gin.Context) {
  var user model.User
  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  
  hashedPassword, err := utils.HashPassword(user.Password)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": "Failed to hash Password",
    })
    return
  }
  user.Password = string(hashedPassword)

  if err := configs.DB.Create(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": "Failed to create user",
    })
    return
  }

  c.JSON(http.StatusCreated, gin.H{
    "message": "User berhasil terdaftar",
  })
}
