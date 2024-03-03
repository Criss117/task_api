package user

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	mysql "unicauca.edu.co/cristian/task-api/src/db"
	user_entity "unicauca.edu.co/cristian/task-api/src/user/entities"
)

func GetUserService(w http.ResponseWriter){
	var users []user_entity.User 
	mysql.DB.Find(&users)
	
	json.NewEncoder(w).Encode(&users)
}
	
func GetUserByIdService(w http.ResponseWriter, id string) {
	var user user_entity.User
	mysql.DB.First(&user, id)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

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

func DeleteUserService(w http.ResponseWriter, id string) {
	var user user_entity.User
	mysql.DB.First(&user, id)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	
	//mysql.DB.Unscoped().Delete(&user)
	mysql.DB.Delete(&user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted"))
}