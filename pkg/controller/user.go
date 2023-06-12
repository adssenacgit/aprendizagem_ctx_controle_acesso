package controller

import (
	"fmt"
	"net/http"

	request "github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/request"
	response "github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/response"
	usecase "github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/usecases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase usecase.UserUseCase
}

func NewUserController() UserController {
	return UserController{}
}

func (uc *UserController) Get(w http.ResponseWriter, r *http.Request) {
	response.StatusMethodNotAllowed(w, r)

}

func (uc *UserController) Post(c *gin.Context) {
	var userBody request.UserBody
	if err := c.BindJSON(&userBody); err != nil {
		fmt.Println(err)
	}

	token, err := uc.UserUseCase.Login(userBody)
	if err != nil {
		panic(err)
	}

	payload := struct {
		Token string `json:"token"`
	}{
		Token: *token,
	}

	c.JSON(200, payload)

}

func (uc *UserController) Put(w http.ResponseWriter, r *http.Request) {
	response.StatusMethodNotAllowed(w, r)

}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	response.StatusMethodNotAllowed(w, r)
}
