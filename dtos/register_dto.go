package dtos

type RegisterDto struct {
	Email    string `json:"email" checkers:"trim email required"`
	Password string `json:"password" checkers:"trim required min-len:8"`
	Username string `json:"username" checkers:"trim required alphanumeric"`
	Country  string `json:"country" checkers:"trim required"`
}
