package routes

import (
	articleRoutes "blog/internal/modules/article/routes"
	homeRoutes "blog/internal/modules/home/routes"
	userRoutes "blog/internal/modules/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)
}
