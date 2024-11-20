package public

import (
	controllers "go-server/app/Controllers"
	public "go-server/app/Controllers/Public"
	middlewares "go-server/app/Middlewares"

	"github.com/gin-gonic/gin"
)

func InitAuthRoute(router *gin.RouterGroup) {
	publicGroup := router.Group("/public")
	{
		publicGroup.POST("/user", controllers.UserCreate)
		publicGroup.POST("/register", public.Register)
		publicGroup.POST("/login", public.Login)
		publicGroup.GET("/profile", middlewares.Auth, public.Profile)
		publicGroup.POST("/forgot-password", public.ForgotPassword)
	}
}
