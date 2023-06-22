package controller

import (
	"fmt"
	"net/http"

	request "github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/request"
	response "github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/response"
	service "github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/service"
	usecase "github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/usecases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase usecase.UserUseCase
}

func NewUserController() UserController {
	return UserController{}
}

func (uc *UserController) Get(c *gin.Context) {
	ucService := service.NewUserService()
	var request struct {
		Id string
	}
	if err := c.BindJSON(&request); err != nil {
		fmt.Println(err)
	}

	user := ucService.GetUserById(request.Id)

	c.JSON(200, user)

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

	c.JSON(200, token)

}

func (uc *UserController) Put(w http.ResponseWriter, r *http.Request) {
	response.StatusMethodNotAllowed(w, r)

}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	response.StatusMethodNotAllowed(w, r)
}
