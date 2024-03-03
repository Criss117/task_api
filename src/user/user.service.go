package user

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	mysql "unicauca.edu.co/cristian/task-api/src/db"
	user_entity "unicauca.edu.co/cristian/task-api/src/user/entities"
)

func CreateUserService(w http.ResponseWriter,user user_entity.User) {
	if user.Name == "" || user.Surname == "" || user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("All fields are required"))
		return
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	user.Password = string(passHash)

	createdUser := mysql.DB.Create(&user)
	
	if createdUser.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createdUser.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}