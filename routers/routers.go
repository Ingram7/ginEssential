package routers

import (
	"ginessential/controller"
	"ginessential/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	gin.ForceConsoleColor()
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
