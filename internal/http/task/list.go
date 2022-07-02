package task

import (
	"encoding/json"
	"gitlab.com/g6834/team41/tasks/internal/domain"
	"net/http"
)

type ListRequest struct {
	Email string `json:"email"`
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: Validate access rights

	// Parse request body
	var req ListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "{}", http.StatusBadRequest)
		return
	}

	// Get all tasks
	tasks, err := domain.ListTasks(req.Email)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}
}
