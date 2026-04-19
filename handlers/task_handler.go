package handlers

import (
	"encoding/json"
	"net/http"
	"taskmanager/models"
	"taskmanager/store"

	"github.com/gorilla/mux"
)
// Get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(store.Tasks)
}

// Create task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	store.Tasks = append(store.Tasks, task)
	json.NewEncoder(w).Encode(task)
}

// Get task by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range store.Tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

// Update task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range store.Tasks {
		if item.ID == params["id"] {
			store.Tasks = append(store.Tasks[:index], store.Tasks[index+1:]...)

			var updated models.Task
			json.NewDecoder(r.Body).Decode(&updated)

			updated.ID = params["id"]
			store.Tasks = append(store.Tasks, updated)

			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

// Delete task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range store.Tasks {
		if item.ID == params["id"] {
			store.Tasks = append(store.Tasks[:index], store.Tasks[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Deleted"})
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}
