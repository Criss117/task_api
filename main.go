package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"unicauca.edu.co/cristian/task-api/src/auth"
	auth_middleware "unicauca.edu.co/cristian/task-api/src/auth/middleware"
	mysql "unicauca.edu.co/cristian/task-api/src/db"
	"unicauca.edu.co/cristian/task-api/src/task"
	task_entity "unicauca.edu.co/cristian/task-api/src/task/entities"
	"unicauca.edu.co/cristian/task-api/src/user"
	user_entity "unicauca.edu.co/cristian/task-api/src/user/entities"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	mysql.DBConnection()
	mysql.DB.AutoMigrate(&user_entity.User{}, &task_entity.Task{})

	var router = mux.NewRouter()

	//public
	router.HandleFunc("/api/user/login", auth.LoginController).Methods("POST")
	router.HandleFunc("/api/user", user.CreateUserController).Methods("POST")

	taskRoutes := mux.NewRouter().PathPrefix("/api/task").Subrouter()
	taskRoutes.Use(auth_middleware.LoggingMiddleware)

	taskRoutes.HandleFunc("", task.CreateTaskController).Methods("POST")
	taskRoutes.HandleFunc("", task.FindAllTaskController).Methods("GET")
	taskRoutes.HandleFunc("/{id}", task.FindTaskController).Methods("GET")
	taskRoutes.HandleFunc("/{id}", task.DeleteTaskController).Methods("DELETE")
	taskRoutes.HandleFunc("/{id}", task.UpdateTaskController).Methods("PUT")

	router.PathPrefix("/api/task").Handler(taskRoutes)

	// r.HandleFunc("/api/tasks", task.FindAllTaskController).Methods("GET")
	// r.HandleFunc("/api/tasks", task.FindTaskController).Methods("GET")
	// r.HandleFunc("/api/tasks", task.FindAllTaskController).Methods("GET")
	// r.HandleFunc("/api/tasks", task.FindAllTaskController).Methods("GET")

	http.ListenAndServe(":8080", router)
}