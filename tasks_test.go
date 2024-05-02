package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateTask(t *testing.T) {
	ms:=&MockStore{}
	service:=NewTasksService(ms)

	t.Run("should return an error if name is empty", func(t *testing.T) {
		payload:=&Task{Name: ""}

		b,err := json.Marshal(payload)

		if err != nil {
			t.Fatal(err)
		}
		req,err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		rr:=httptest.NewRecorder()
		router:=mux.NewRouter()

		router.HandleFunc("/tasks", service.handleCreateTask) 
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
		}

	})


}

func TestGetTask(t *testing.T) {
	ms:=&MockStore{}
	service:=NewTasksService(ms)

	t.Run("should return an error if task is not found", func(t *testing.T) {
		req,err:=http.NewRequest(http.MethodGet, "/tasks/42", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr:=httptest.NewRecorder()
		router:=mux.NewRouter()

		router.HandleFunc("/tasks/{id}", service.handleGetTasks)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, rr.Code)
		}
	})
}