package task

import (
	"net/http"

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
		return
	}

	// Try to parse.
	id, ok := rawId.(string)
	if !ok {
		http.Error(w, "{}", http.StatusBadRequest)
		return
	}

	err := domain.DeleteTask(id)
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
