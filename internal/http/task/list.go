package task

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"gitlab.com/g6834/team41/tasks/internal/domain"
)

type ListRequest struct {
	Email string `json:"email" example:"test@test.org"`
}

// @Summary List of tasks
// @Description List of tasks
// @Accept json
// @Produce json
// @Param task body ListRequest true "List"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /task/ [get]
func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log := logrus.WithField("RAPI", "List tasks")

	// Parse request body
	var req ListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error("error decode request body: ", err)
		http.Error(w, "{}", http.StatusBadRequest)
		return
	}

	// Get all tasks
	tasks, err := domain.ListTasks(req.Email)
	if err != nil {
		log.Error("domain.ListTasks error: ", err)
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(tasks)
	if err != nil {
		log.Error("error marshaling response: ", err)
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		log.Error("error writing response: ", err)
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}
}
