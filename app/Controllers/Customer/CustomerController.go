package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Customer Controller
func Customer(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User profile", "user": user})
}
