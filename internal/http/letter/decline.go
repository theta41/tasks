package letter

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/domain"
	"gitlab.com/g6834/team41/tasks/internal/http/middlewares"
)

func Decline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log := logrus.WithField("RAPI", "Decline letter by uuid")

	// Get uuid from /decline/{uuid}
	rawUuid := r.Context().Value(middlewares.ContextKeyTaskUuid)
	if rawUuid == nil {
		http.Error(w, "{}", http.StatusBadRequest)
		log.Error("missing letter uuid in context")
		return
	}

	// Try to parse.
	strUuid, ok := rawUuid.(string)
	if !ok {
		http.Error(w, "{}", http.StatusBadRequest)
		log.Error("cannot cast letter uuid to string: ", rawUuid)
		return
	}

	uid, err := uuid.Parse(strUuid)
	if err != nil {
		http.Error(w, "{}", http.StatusBadRequest)
		log.Error("cannot parse letter uuid: ", strUuid)
		return
	}

	err = domain.Decline(uid)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		log.Error("domain.Decline error: ", err)
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
