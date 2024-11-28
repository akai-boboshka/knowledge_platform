package middleware

import (
	"awesomeProject/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")
		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}

		if len(authToken) > 7 && authToken[:7] == "Bearer " {
			authToken = authToken[7:]
		}

		userID, roleID, err := utils.ValidateJWT(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}

		c.Set("id", userID)
		c.Set("role", roleID)
		c.Next()
	}
}
