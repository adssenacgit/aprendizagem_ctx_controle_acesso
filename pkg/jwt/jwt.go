package jwt

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
