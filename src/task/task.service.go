package task

import (
	"encoding/json"
	"net/http"

	mysql "unicauca.edu.co/cristian/task-api/src/db"
	task_entity "unicauca.edu.co/cristian/task-api/src/task/entities"
)

func CreateTaskService(w http.ResponseWriter, task task_entity.Task) {
	if task.Title == "" || task.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("All fields are required"))
		return
	}
	taskCreated := mysql.DB.Create(&task)

	if taskCreated.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(taskCreated.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func FindAllTaskService(w http.ResponseWriter, userId uint) {
	var tasks []task_entity.Task 
	mysql.DB.Find(&tasks, "user_id = ?", userId)
	
	json.NewEncoder(w).Encode(&tasks)
}

func FindTaskService(w http.ResponseWriter, taskId string, userId uint) {
	var task task_entity.Task
	mysql.DB.First(&task, "id = ? AND user_id = ?", taskId, userId)

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}


func DeleteTaskService(w http.ResponseWriter, taskId string, userId uint) {
	var task task_entity.Task
	mysql.DB.First(&task, "id = ? AND user_id = ?", taskId, userId)

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	mysql.DB.Delete(&task)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task deleted"))
}

func UpdateTaskService(w http.ResponseWriter, task task_entity.Task, userId uint) {
	var oldTask task_entity.Task
	mysql.DB.First(&oldTask, "id = ? AND user_id = ?", task.ID, userId)

	if oldTask.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	
	mysql.DB.Save(&task)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task updated"))
}