package routes

import (
	controller "github.com/folafunmi-db/go-jwt-auth/controllers"
	"github.com/folafunmi-db/go-jwt-auth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
