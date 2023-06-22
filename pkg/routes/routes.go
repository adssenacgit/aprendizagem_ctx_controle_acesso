package routes

import (
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Router *gin.Engine

func CreateUrlMappings(DbConn *gorm.DB) {
	Router = gin.Default()

	userController := controller.NewUserController()
	jwtController := controller.NewJwtController()

	v1 := Router.Group("/v1")
	{
		v1.POST("/login", userController.Post)
		v1.GET("/keys", jwtController.GetKeys)
		v1.GET("/user", userController.Get)
	}

}
