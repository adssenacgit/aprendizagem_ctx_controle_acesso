package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/entities"
	"github.com/golang-jwt/jwt/v5"
)

type KeyData struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
	Alg string `json:"alg"`
	K   string `json:"k"`
	Typ string `json:"typ"`
}

type PublicKeysData struct {
	Keys []KeyData `json:"keys"`
}

func GenerateJWT(userData entities.User) string {
	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	privateKeyBytes, err := os.ReadFile(fmt.Sprintf("%s/certs/private.pem", currPath))
	if err != nil {
		panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	jti := make([]byte, 16)
	_, err = rand.Read(jti)
	if err != nil {
		panic(err)
	}

	var l [2]string
	l[0] = "a"
	l[1] = "b"

	now := time.Now()
	expirationDate := now.Add(time.Hour * 3)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256,

		jwt.MapClaims{
			"sub":   userData.Id,
			"kid":   "go-ext-authz",
			"roles": l,
			"jti":   hex.EncodeToString(jti),
			"aud":   "https://your.krakend.io",
			"iss":   "https://your-backend", // emissor
			"iat":   now.Unix(),
			"exp":   expirationDate.Unix(),
		})
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	return signedToken
}

func GetJWK() PublicKeysData {

	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	publicKeyBytes, err := os.ReadFile(fmt.Sprintf("%s/certs/public.pem", currPath))
	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		panic(err)
	}

	n := base64.StdEncoding.EncodeToString((*publicKey.N).Bytes())

	e := base64.StdEncoding.EncodeToString(big.NewInt(int64(publicKey.E)).Bytes())

	return PublicKeysData{
		Keys: []KeyData{{"RSA", "go-ext-authz", "sig", n, e, "RS256", n, "JWT"}},
	}

}
