package usecases

import (
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/entities"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/jwt"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/request"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/service"
)

type UserUseCase struct {
	service service.UserService
}

func NewUserUseCase(service service.UserService) UserUseCase {
	return UserUseCase{}
}

func (uuc *UserUseCase) Login(userData request.UserBody) (*string, error) {
	var user entities.User
	if userData.Username == "admin" && userData.Password == "admin" {
		user.Id = "1234"
		user.Password = "admin"
		user.Roles = append(user.Roles, "a")
		user.Roles = append(user.Roles, "b")
		token := jwt.GenerateJWT(user)
		return &token, nil
	}

	return nil, nil
}

func (uuc *UserUseCase) CreateUser(request.UserBody) {
}

func (uuc *UserUseCase) GetJWK() jwt.PublicKeysData {
	return jwt.GetJWK()
}
