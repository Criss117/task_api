package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	user_entitie "unicauca.edu.co/cristian/task-api/src/user/entities"
)

func CreateUserController(w http.ResponseWriter, r *http.Request) {
	var user user_entitie.User
	json.NewDecoder(r.Body).Decode(&user)
	CreateUserService(w, user)
}

func FindUsersController(w http.ResponseWriter, r *http.Request) {
	GetUserService(w)
}

func FindUserController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	GetUserByIdService(w, userId)
}


func DeleteUserController(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userId := params["id"]
		DeleteUserService(w, userId)
}