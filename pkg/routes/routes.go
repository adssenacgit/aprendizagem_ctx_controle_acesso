package routes

import (
	"database/sql"

	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/controller"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings(DbConn *sql.DB) {
	Router = gin.Default()

	userController := controller.NewUserController()
	jwtController := controller.NewJwtController()

	v1 := Router.Group("/v1")
	{
		v1.POST("/login", userController.Post)
		v1.GET("/.well-known/jwks", jwtController.Get)
	}

}
