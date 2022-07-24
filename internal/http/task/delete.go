package task

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/domain"
)

// @Summary Delete task
// @Description Delete task
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /task/{id}/ [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get id from /task/{id}
	rawId := r.Context().Value("id")
	if rawId == nil {
		http.Error(w, "{}", http.StatusBadRequest)
		logrus.Error("missing task id in context")
		return
	}

	// Try to parse.
	id, ok := rawId.(string)
	if !ok {
		http.Error(w, "{}", http.StatusBadRequest)
		logrus.Error("can not cast task id to string: ", rawId)
		return
	}

	err := domain.DeleteTask(id)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		logrus.Error("domain.DeleteTask error: ", err)
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
