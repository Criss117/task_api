package task

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	auth_middleware "unicauca.edu.co/cristian/task-api/src/auth/middleware"
	task_entity "unicauca.edu.co/cristian/task-api/src/task/entities"
	user_entity "unicauca.edu.co/cristian/task-api/src/user/entities"
)

func FindAllTaskController(w http.ResponseWriter, r *http.Request) {
	loggedUser := r.Context().Value(auth_middleware.LoggedInUser).(user_entity.User)
	FindAllTaskService(w, loggedUser.ID)
}

func CreateTaskController(w http.ResponseWriter, r *http.Request) {
	var task task_entity.Task
	json.NewDecoder(r.Body).Decode(&task)

	loggedUser := r.Context().Value(auth_middleware.LoggedInUser).(user_entity.User)
	task.UserID = loggedUser.ID
	CreateTaskService(w, task)
}

func FindTaskController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId := params["id"]
	loggedUser := r.Context().Value(auth_middleware.LoggedInUser).(user_entity.User)

	FindTaskService(w, taskId, loggedUser.ID)
}

func DeleteTaskController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId := params["id"]
	loggedUser := r.Context().Value(auth_middleware.LoggedInUser).(user_entity.User)

	DeleteTaskService(w, taskId, loggedUser.ID)
}