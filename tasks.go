package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)
var errNameRequired = errors.New("name is required")
var errProjectIDRequired = errors.New("projectID is required")
var errUserIDRequired = errors.New("userID is required")

type TasksService struct {
	store Store
}

func NewTasksService(store Store) *TasksService {
	return &TasksService{store: store}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", WithJWTAuth(s.handleCreateTask,s.store)).Methods("POST")
	r.HandleFunc("/tasks/{id}",WithJWTAuth(s.handleGetTasks,s.store) ).Methods("GET")
}

func (s *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	// code
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// handle error
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}
	defer r.Body.Close()
	var task *Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		// handle error
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}
	if err=validateTaskPayload(task) ; err != nil {
		// handle error
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	t,err:=s.store.CreateTask(task)
	if err != nil {
		// handle error
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating task"})
		return
	}
	WriteJSON(w, http.StatusCreated, t)

}

func (s *TasksService) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	// code
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "id is required"})
		return
	}
	t, err := s.store.GetTask(id)
	
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error getting task"})
		return
	}
	WriteJSON(w, http.StatusOK, t)
}
func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errNameRequired
	}

	if task.ProjectID == 0 {
		return errProjectIDRequired
	}

	if task.AssignedToID == 0 {
		return errUserIDRequired
	}

	return nil
}