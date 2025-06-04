package dtos

type LoginDto struct {
	Email    string `json:"email" checkers:"trim"`
	Password string `json:"password" checkers:"trim required"`
	Username string `json:"username" checkers:"trim"`
}
