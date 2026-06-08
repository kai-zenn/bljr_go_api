package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kai-zenn/bljr_go_api/api/utils"
)

func RBACMiddleware(permission string) gin.HandlerFunc {
  return func(c *gin.Context) {
    id, exists := c.Get("id")
    if !exists {
      c.JSON(http.StatusUnauthorized, gin.H{
        "error": "User ID not found",
      })
      c.Abort()
      return
    }
    userId, err := uuid.Parse(id.(string))
    if err != nil {
      c.JSON(http.StatusUnauthorized, gin.H{
        "error": "User ID not found",
      })
      c.Abort()
      return
    }
    
    if !utils.HasAccess(userId, permission) {
      c.JSON(http.StatusForbidden, gin.H{
        "error": "Access denied",
      })
      c.Abort()
      return
    }
    c.Next()
  }
}
