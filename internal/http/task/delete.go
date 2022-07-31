package task

import (
	"gitlab.com/g6834/team41/tasks/internal/env"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/domain"
	"gitlab.com/g6834/team41/tasks/internal/http/middlewares"
)

// @Summary Delete task
// @Description Delete task
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /tasks/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log := logrus.WithField("RAPI", "Delete task by id")

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

	err := domain.DeleteTask(id)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		log.Error("domain.DeleteTask error: ", err)
		return
	}

	// Send task to kafka
	err = env.E.K.Publish([]byte("delete"), []byte(strconv.Itoa(id)))
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		logrus.Error("env.E.K.Publish error: ", err)
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
