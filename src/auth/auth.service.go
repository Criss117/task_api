package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	login_dto "unicauca.edu.co/cristian/task-api/src/auth/dto"
	mysql "unicauca.edu.co/cristian/task-api/src/db"
	user_entity "unicauca.edu.co/cristian/task-api/src/user/entities"
)

type response struct {
	JWT string `json:"token"`
}	

func login(w http.ResponseWriter, loginUserDto login_dto.LoginUserDto) {
	var user user_entity.User
	mysql.DB.Find(&user, "email = ?", loginUserDto.Email)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUserDto.Password));

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid credentials"))
		return
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := jwt.SignedString([]byte(os.Getenv("SECRETKEY")))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	res := response{
		JWT: token,
	}

	json.NewEncoder(w).Encode(res)
}