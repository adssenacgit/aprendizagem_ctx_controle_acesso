package controller

import (
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type JwtController struct {
}

func NewJwtController() JwtController {
	return JwtController{}
}

func (jc *JwtController) GetKeys(c *gin.Context) {
	c.JSON(200, jwt.GetKeys())
}
