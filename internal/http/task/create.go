package task

import (
	"encoding/json"
	"gitlab.com/g6834/team41/tasks/internal/domain"
	"gitlab.com/g6834/team41/tasks/internal/models"
	"net/http"
	"time"
)

type CreateRequest struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	CreatorEmail      string   `json:"creator_email"`
	CreatedAt         int      `json:"created_at"`
	FinishedAt        int      `json:"finished_at"`
	ParticipantEmails []string `json:"participant_emails"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "{}", http.StatusBadRequest)
		return
	}

	// TODO: Validate access rights

	// Create task
	err := domain.CreateTask(models.Task{
		Name:         req.Name,
		Description:  req.Description,
		CreatorEmail: req.CreatorEmail,
		CreatedAt:    time.Unix(int64(req.CreatedAt), 0),
		FinishedAt:   time.Unix(int64(req.FinishedAt), 0),
	}, req.ParticipantEmails)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}
}
