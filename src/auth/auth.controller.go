package auth

import (
	"encoding/json"
	"net/http"

	login_dto "unicauca.edu.co/cristian/task-api/src/auth/dto"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	var user login_dto.LoginUserDto
	json.NewDecoder(r.Body).Decode(&user)
	login(w, user)
}