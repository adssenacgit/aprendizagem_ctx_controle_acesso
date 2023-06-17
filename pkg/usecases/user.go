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

func (uuc *UserUseCase) Login(userData request.UserBody) (interface{}, error) {
	var user entities.User
	if userData.Username == "admin" && userData.Password == "admin" {
		user.Id = "1234"
		user.Password = "admin"
		user.Roles = append(user.Roles, "a")
		user.Roles = append(user.Roles, "b")
		// token := jwt.GenerateJWT(user)

		type Token struct {
			Aud   string   `json:"aud"`
			Iss   string   `json:"iss"`
			Sub   string   `json:"sub"`
			Jti   string   `json:"jti"`
			Roles []string `json:"roles"`
			Exp   int64
		}

		type Payload struct {
			AccessToken  Token `json:"access_token"`
			RefreshToken Token `json:"refresh_token"`
		}

		accessToken := Token{
			Aud:   "https://your.krakend.io",
			Iss:   "https://your-backend",
			Sub:   "1234567890qwertyuio",
			Jti:   "mnb23vcsrt756yuiomnbvcx98ertyuiop",
			Roles: user.Roles,
			Exp:   1735689600,
		}

		refresh_token := Token{
			Aud: "https://your.krakend.io",
			Iss: "https://your-backend",
			Sub: "1234567890qwertyuio",
			Jti: "mnb23vcsrt756yuiomn12876bvcx98ertyuiop",
			Exp: 1735689600,
		}

		payload := Payload{
			AccessToken:  accessToken,
			RefreshToken: refresh_token,
		}

		return payload, nil
	}
	return nil, nil
}

func (uuc *UserUseCase) CreateUser(request.UserBody) {
}

func (uuc *UserUseCase) GetJWK() jwt.PublicKeysData {
	return jwt.GetJWK()
}
