package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kai-zenn/bljr_go_api/api/utils"
)


func AuthMiddleware() gin.HandlerFunc{
  return func(c *gin.Context) {
    tokenHeader := c.GetHeader("Authorization")
    if tokenHeader == "" || !strings.HasPrefix(tokenHeader, "Bearer") {
      c.JSON(http.StatusUnauthorized, gin.H{
        "error": "Unauthorized access",
      })
      c.Abort()
      return
    }

    tokenStr := strings.TrimPrefix(tokenHeader, "Bearer ")

    token, err := utils.VerifyToken(tokenStr)
    if err != nil || !token.Valid {
      c.JSON(http.StatusUnauthorized, gin.H{
        "error": "Invalid or Expired Token",
      })
      c.Abort()
      return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userID := fmt.Sprintf("%v", claims["user_id"])

		c.Set("id", userID)
    
    c.Next()
  }
}
