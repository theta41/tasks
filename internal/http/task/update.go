package task

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/domain"
	"gitlab.com/g6834/team41/tasks/internal/http/middlewares"
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
// @Router /tasks/{id} [put]
func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log := logrus.WithField("RAPI", "Update task with id")

	// Get id from /tasks/{id}
	rawId := r.Context().Value(middlewares.ContextKeyTaskId)
	if rawId == nil {
		http.Error(w, "{}", http.StatusBadRequest)
		log.Error("missing task id in context")
		return
	}

	// Try to parse.
	id, ok := rawId.(int)
	if !ok {
		http.Error(w, "{}", http.StatusBadRequest)
		log.Error("can not cast task id to int: ", rawId)
		return
	}

	// Parse body.
	var req models.Task
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "{}", http.StatusBadRequest)
		log.Error("error decode request body: ", err)
		return
	}

	if req.ID != id {
		http.Error(w, "{}", http.StatusBadRequest)
		log.Errorf("wrong task id %v for task %v", id, req)
		return
	}

	// Get task
	err = domain.UpdateTask(req)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		log.Error("domain.UpdateTask error: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		log.Error("error write response: ", err)
		return
	}
}
