package letter

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/env"
	"net/http"

	"github.com/google/uuid"
	"gitlab.com/g6834/team41/tasks/internal/domain"
)

func Accept(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//TODO: Validate access rights.

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

	err = domain.Accept(uid)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}

	// Send letter to kafka
	err = env.E.K.Publish([]byte("accept"), []byte(id))
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		logrus.Error("env.E.K.Publish error: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		return
	}
}
