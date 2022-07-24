package task

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/domain"
	"gitlab.com/g6834/team41/tasks/internal/models"
)

// @Summary Update task
// @Description Update task
// @Accept json
// @Produce json
// @Param task body models.Task true "Task"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /task/{id}/ [put]
func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get id from /task/{id}
	rawId := r.Context().Value("id")
	if rawId == nil {
		http.Error(w, "{}", http.StatusBadRequest)
		logrus.Error("missing task id in context")
		return
	}

	// Parse body.
	var req models.Task
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "{}", http.StatusBadRequest)
		logrus.Error("error decode request body: ", err)
		return
	}

	// Get task
	err = domain.UpdateTask(req)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		logrus.Error("domain.UpdateTask error: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		logrus.Error("error write response: ", err)
		return
	}
}
