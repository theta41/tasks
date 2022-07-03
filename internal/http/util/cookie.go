package util

import (
	"net/http"

	"gitlab.com/g6834/team41/tasks/internal/models"
)

func GetTokensFromCookie(r *http.Request) models.TokenPair {
	return models.TokenPair{
		AccessToken:  "-AccessToken-",
		RefreshToken: "-RefreshToken-",
	}
}

func PutTokensToCookie(w http.ResponseWriter, tokens models.TokenPair) {
	// TODO Put new tokens to cookie
}
