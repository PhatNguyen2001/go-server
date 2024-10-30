package public

import (
	controllers "go-server/app/Controllers"
	customer "go-server/app/Controllers/Customer"
	public "go-server/app/Controllers/Public"
	middlewares "go-server/app/Middlewares"

	"github.com/gin-gonic/gin"
)

func InitAuthRoute(router *gin.RouterGroup) {
	publicGroup := router.Group("/public")
	{
		publicGroup.POST("/user", controllers.UserCreate)
		// publicGroup.POST("/register", public.Register)
		publicGroup.POST("/login", public.Login)
		publicGroup.GET("/customer", middlewares.Auth, customer.Customer)
	}
}
