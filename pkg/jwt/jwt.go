package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
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

type PrivateKeyData struct {
	Kty string
	Crv string
	X   string
	Y   string
	Kid string
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

	publicKeyBytes, err := os.ReadFile(fmt.Sprintf("%s/certs/public_key.pem", currPath))
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

func GetPrivateKey() PrivateKeyData {
	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	privateJwk, err := os.ReadFile(fmt.Sprintf("%s/certs/private_key.jwk", currPath))
	if err != nil {
		panic(err)
	}

	var jwkKey PrivateKeyData

	json.Unmarshal(privateJwk, &jwkKey)

	return jwkKey
}

func GetKeys() interface{} {
	type Key struct {
		Kty string `json:"kty"`
		Alg string `json:"alg"`
		K   string `json:"k"`
		Kid string `json:"kid"`
	}

	type Keys struct {
		Keys []Key `json:"keys"`
	}

	key1 := Key{
		Kty: "oct",
		Alg: "A128KW",
		K:   "GawgguFyGrWKav7AX4VKUg",
		Kid: "sim1",
	}

	key2 := Key{
		Kty: "oct",
		K:   "AyM1SysPpbyDfgZld3umj1qzKObwVMkoqQ-EstJQLr_T-1qS0gZH75aKtMN3Yj0iPS4hcgUuTwjAzZr1Z9CAow",
		Kid: "sim2",
		Alg: "HS256",
	}

	return Keys{
		Keys: []Key{key1, key2},
	}

}
