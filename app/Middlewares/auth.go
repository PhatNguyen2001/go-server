package middlewares

import (
	"fmt"
	"go-server/database/connection"
	"go-server/database/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(c *gin.Context) {
	fmt.Print("In middleWare")
	authHeader := c.GetHeader("Authorization")

	fmt.Printf("Authorization Header: '%s'\n", authHeader)

	//Null or not have Prefix Bearer + token
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
		c.Abort()
		return
	}

	//replac Prefix Bearer
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	//Parse Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find User
		var user models.User
		connection.DB.First(&user, "id=?", claims["id"])

		if user.Id == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//attach to request
		c.Set("user", user)

		fmt.Print(claims["id"])
	} else {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
