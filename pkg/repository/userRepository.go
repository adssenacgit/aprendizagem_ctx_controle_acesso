package repository

import (
	"context"
	"database/sql"

	entities "github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/entities"
)

type UserRepository struct {
	dbConn *sql.DB
}

func NewUserRepository(dbConection *sql.DB) UserRepository {
	return UserRepository{
		dbConn: dbConection,
	}
}

func (ur *UserRepository) GetUserByUsername(username string) (*entities.User, error) {
	user := &entities.User{}

	row, err := ur.dbConn.Query("SELECT * FROM users WHERE username=?", username)
	if err != nil {
		return nil, err
	}

	if err := row.Scan(&user.Id, &user.Roles, &user.Password); err != nil {
		return nil, err
	}

	return user, nil

}

func (ur *UserRepository) GetUserById(id string) (*entities.User, error) {

	user := &entities.User{}

	row, err := ur.dbConn.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	if err := row.Scan(&user.Id, &user.Roles, &user.Password); err != nil {
		return nil, err
	}

	return user, nil

}

func (ur *UserRepository) CreateUser(id string, name string, hashedPassword string) error {
	query := "INSERT INTO users (id, name, password) VALUES (?, ?, ?)"
	_, err := ur.dbConn.ExecContext(context.Background(), query, id, name, hashedPassword)
	if err != nil {
		return err
	}

	return nil

}
