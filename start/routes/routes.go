package routes

import (
	"go-server/start/routes/public"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api/v1")
	public.InitAuthRoute(apiGroup)

	return r
}
