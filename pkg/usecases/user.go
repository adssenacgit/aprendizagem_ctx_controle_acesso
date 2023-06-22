package usecases

import (
	"time"

	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/request"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/service"
)

type UserUseCase struct {
	service service.UserService
}

type Token struct {
	Aud    string `json:"aud"`
	Iss    string `json:"iss"`
	Sub    string `json:"sub"`
	Jti    string `json:"jti"`
	Role   string `json:"role"`
	UserId string `json:"user_id"`
	Exp    int64
}

type Payload struct {
	AccessToken  Token `json:"access_token"`
	RefreshToken Token `json:"refresh_token"`
}

func NewUserUseCase(service service.UserService) UserUseCase {
	return UserUseCase{}
}

func (uuc *UserUseCase) Login(userData request.UserBody) (interface{}, error) {

	user := uuc.service.GetUserById(userData.Login)
	if user == nil {
		panic("usuario não existe")
	}

	// checkPw := cripto.CheckPasswordHash(userData.Password, user.Password)
	// if !checkPw {
	// 	panic("senha não bate")
	// }

	now := time.Now()

	accessToken := Token{
		Aud:    "https://your.krakend.io",
		Iss:    "https://your-backend",
		Sub:    "1234567890qwertyuio",
		Jti:    "mnb23vcsrt756yuiomnbvcx98ertyuiop",
		Role:   user.Role,
		UserId: user.Id.String(),
		Exp:    now.Add(time.Minute * 15).Unix(),
	}

	refresh_token := Token{
		Aud:    "https://your.krakend.io",
		Iss:    "https://your-backend",
		Sub:    "1234567890qwertyuio",
		Role:   user.Role,
		UserId: user.Id.String(),
		Jti:    "mnb23vcsrt756yuiomn12876bvcx98ertyuiop",
		Exp:    now.Add(time.Hour * 168).Unix(),
	}

	payload := Payload{
		AccessToken:  accessToken,
		RefreshToken: refresh_token,
	}

	return payload, nil
}

func (uuc *UserUseCase) CreateUser(request.UserBody) {
}
