package repository

import (
	entities "github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/entities"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/infrastructure"
	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository(dbConection *gorm.DB) UserRepository {
	return UserRepository{}
}

func (ur *UserRepository) GetUserById(id string) *entities.User {

	user := &entities.User{}

	infrastructure.DB.First(user, "id = ?", id)

	return user

}

// func (ur *UserRepository) CreateUser(id string, name string, hashedPassword string) error {

// }
