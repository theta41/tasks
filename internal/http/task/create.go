package task

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/domain"
	"gitlab.com/g6834/team41/tasks/internal/env"
	"gitlab.com/g6834/team41/tasks/internal/http/util"
	"gitlab.com/g6834/team41/tasks/internal/models"
)

type CreateRequest struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	ParticipantEmails []string `json:"participant_emails"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "{}", http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	// TODO: Get login from cookie
	login := "test@example.org"

	tokens := util.GetTokensFromCookie(r)

	// Validate access rights
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	newTokens, err := env.E.Auth.Validate(ctx, login, tokens)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	util.PutTokensToCookie(w, newTokens)

	// Create task
	err = domain.CreateTask(models.Task{
		Name:        req.Name,
		Description: req.Description,
		// TODO: change
		CreatorEmail: login,
		CreatedAt:    time.Now(),
		FinishedAt:   time.Unix(int64(0), 0),
	}, req.ParticipantEmails)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	// Send response
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
		logrus.Error(err)
		return
	}
}
