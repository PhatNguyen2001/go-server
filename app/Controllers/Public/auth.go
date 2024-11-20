package public

import (
	"fmt"
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
		Gender   string
		Bod      string
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

	customer := models.Customer{
		Gender: body.Gender,
		Bod:    body.Bod,
		UserId: user.Id,
	}
	customerResult := connection.DB.Create(&customer)
	if customerResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Failed to create customer",
		})
		return
	}

	connection.DB.Preload("User").First(&customer, "id = ?", customer.Id)
	c.JSON(http.StatusOK, gin.H{
		"customer": customer,
	})
}

func Login(c *gin.Context) {
	var body struct {
		Password string
		Phone    string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Failed to read body!",
		})
		return
	}

	var user models.User
	connection.DB.First(&user, "phone=?", body.Phone)

	fmt.Printf("User: %+v\n", user)
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

func Profile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User profile", "user": user})
}

func ForgotPassword(c *gin.Context) {
	var body struct {
		Phone string `json:"phone" binding:"required"`
		Email string `json:"email" binding:"required"`
	}

	// Bind input
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	// Find user by phone
	var user models.User
	if err := connection.DB.First(&user, "phone = ? AND email = ?", body.Phone, body.Email).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	// Generate temporary password
	tempPassword := helper.GenerateTemporaryPassword(8)

	// Update user's password (hashed version for security)
	user.Password = tempPassword
	connection.DB.Save(&user)

	// Send email with temporary password
	if err := helper.SendResetEmail(user.Email, tempPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send email",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Temporary password sent to your email",
	})
}
