package task

import (
	"encoding/json"
	"gitlab.com/g6834/team41/tasks/internal/domain"
	"gitlab.com/g6834/team41/tasks/internal/models"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get id from /task/{id}
	rawId := r.Context().Value("id")
	if rawId == nil {
		http.Error(w, "{}", http.StatusBadRequest)
		return
	}

	// Parse body.
	var req models.Task
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "{}", http.StatusBadRequest)
		return
	}

	// Get task
	err = domain.UpdateTask(req)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}
}
