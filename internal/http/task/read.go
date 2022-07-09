package task

import (
	"encoding/json"
	"net/http"

	"gitlab.com/g6834/team41/tasks/internal/domain"
	"gitlab.com/g6834/team41/tasks/internal/models"
)

type ReadRequest struct {
	Name string `json:"name" example:"Test task"`
}

// @Summary Read task
// @Description Read task
// @Accept json
// @Produce json
// @Param task body ReadRequest true "Read"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /task/{id}/ [get]
func Read(w http.ResponseWriter, r *http.Request) {
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

	// Get task
	task, letters, err := domain.GetTask(id)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}

	// Send response
	resp := struct {
		Task    models.Task     `json:"task"`
		Letters []models.Letter `json:"letters"`
	}{
		Task:    *task,
		Letters: letters,
	}

	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(res)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}
}
