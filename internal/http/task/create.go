package task

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/domain"
	"gitlab.com/g6834/team41/tasks/internal/http/util"
	"gitlab.com/g6834/team41/tasks/internal/models"
)

type CreateRequest struct {
	Name              string   `json:"name" example:"Test name"`
	Description       string   `json:"description" example:"Test description"`
	ParticipantEmails []string `json:"participant_emails" swaggertype:"array,string" example:"['test@test.org, test@test.com']"`
}

// @Summary Create task
// @Description Create task
// @Accept json
// @Produce json
// @Param task body CreateRequest true "New Task"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /task/ [post]
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "{}", http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	login := util.GetLoginFromCookie(r)

	// Create task
	err := domain.CreateTask(models.Task{
		Name:         req.Name,
		Description:  req.Description,
		CreatorEmail: login,
		CreatedAt:    time.Now(),
		FinishedAt:   time.Unix(int64(0), 0),
	}, req.ParticipantEmails)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	// Send response
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		logrus.Error(err)
		return
	}
}
