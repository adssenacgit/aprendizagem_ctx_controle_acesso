package service

import (
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/cripto"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/entities"
	repository "github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/repository"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/request"
	"github.com/google/uuid"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return UserService{
		repository: userRepository,
	}
}

func (us *UserService) GetUserById(id string) (*entities.User, error) {
	user, err := us.repository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) CreateUser(userData request.UserBody) error {

	user, err := us.repository.GetUserById(userData.Username)
	if err != nil {
		return err
	}

	if user == nil {
		// TODO: retornar erro
	}

	id := uuid.New().String()
	hashedPassword, err := cripto.HashPassword(userData.Password)
	if err != nil {
		return err
	}

	err = us.repository.CreateUser(id, userData.Username, hashedPassword)
	if err != nil {
		return err
	}
	return nil

}
