package letter

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/domain"
	"gitlab.com/g6834/team41/tasks/internal/env"
	"net/http"
)

func Decline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: Validate access rights.

	// Get id
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

	uid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "{}", http.StatusBadRequest)
		return
	}

	// Send letter to kafka
	err = env.E.K.Publish([]byte("decline"), []byte(id))
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		logrus.Error("env.E.K.Publish error: ", err)
		return
	}

	err = domain.Decline(uid)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}
}
