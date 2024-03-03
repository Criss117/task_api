package user

import (
	"encoding/json"
	"net/http"

	user_entitie "unicauca.edu.co/cristian/task-api/src/user/entities"
)

func CreateUserController(w http.ResponseWriter, r *http.Request) {
	var user user_entitie.User
	json.NewDecoder(r.Body).Decode(&user)
	CreateUserService(w, user)
}