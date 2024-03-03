package login_dto

type LoginUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}