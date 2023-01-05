package model

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type UserRegis struct {
	Username string `validate:"required" json:"username"`
	Name     string `validate:"required" json:"name"`
	Password string `validate:"required" json:"password"`
	Address  string `validate:"required" json:"address"`
	Email    string `validate:"required" json:"email"`
}

type UserUpdate struct {
	Id       int64  `validate:"required" json:"id"`
	Name     string `validate:"required" json:"name"`
	Password string `validate:"required" json:"password"`
	Address  string `validate:"required" json:"address"`
	Email    string `validate:"required" json:"email"`
}

type LoginRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
