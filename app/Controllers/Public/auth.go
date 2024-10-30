package public

import (
	helper "go-server/Utils"
	"go-server/database/connection"
	"go-server/database/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body struct {
		Name     string
		Password string
		Phone    string
		Email    string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Failed to read body!",
		})
		return
	}

	user := models.User{Name: body.Name, Password: body.Password, Phone: body.Phone, Email: body.Email}
	result := connection.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusCreated, gin.H{
			"err": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func Login(c *gin.Context) {
	var body struct {
		// Name     string
		Password string
		Phone    string
		// Email    string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Failed to read body!",
		})
		return
	}

	var user models.User
	connection.DB.First(&user, "phone=?", body.Phone)

	if user.Id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "user not found!",
		})
		return
	}

	// compare Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "user not found!",
		})
		return
	}

	//Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	//Access Token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "inValid create Access Token!",
		})
		return
	}

	//Refresh Token
	refreshToken, err := helper.GenerateRefreshToken(user.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "inValid create refresh token!",
		})
		return
	}

	// c.Header("Authorization", )
	//send it back
	c.JSON(http.StatusOK, gin.H{
		"token":         tokenString,
		"refresh_token": refreshToken,
	})
}
