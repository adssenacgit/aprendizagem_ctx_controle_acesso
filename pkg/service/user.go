package service

import (
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/entities"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/infrastructure"
	repository "github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService() UserService {
	userReposutory := repository.NewUserRepository(infrastructure.DB)
	return UserService{
		repository: userReposutory,
	}
}

func (us *UserService) GetUserById(id string) *entities.User {
	return us.repository.GetUserById(id)
}

// func (us *UserService) CreateUser(userData request.UserBody) error {

// 	user, err := us.repository.GetUserById(userData.Username)
// 	if err != nil {
// 		return err
// 	}

// 	if user == nil {
// 		// TODO: retornar erro
// 	}

// 	id := uuid.New().String()
// 	hashedPassword, err := cripto.HashPassword(userData.Password)
// 	if err != nil {
// 		return err
// 	}

// 	err = us.repository.CreateUser(id, userData.Username, hashedPassword)
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }
