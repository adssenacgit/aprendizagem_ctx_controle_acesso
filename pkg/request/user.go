package request

type UserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
