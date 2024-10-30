package controllers

import (
	"fmt"
	"go-server/database/connection"
	"go-server/database/models"
	"reflect"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	user := models.User{Phone: "0902975726", Email: "caophat225", Name: "Phat", Password: "asdasd123"}

	result := connection.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	fmt.Printf("User: %+v\n", reflect.TypeOf(user))
	fmt.Printf("User: %+v\n", user.Id)

	c.JSON(200, gin.H{
		"user": user,
	})
}
