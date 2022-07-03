package util

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"gitlab.com/g6834/team41/tasks/internal/models"
)

const (
	cookieLogin        = "Login"
	cookieAccessToken  = "AccessToken"
	cookieRefreshToken = "RefreshToken"
)

func GetTokensFromCookie(r *http.Request) models.TokenPair {

	access, _ := r.Cookie(cookieAccessToken)
	refresh, _ := r.Cookie(cookieRefreshToken)

	logrus.Print("got tokens from cookies", access, refresh)

	return models.TokenPair{
		AccessToken:  access.Value,
		RefreshToken: refresh.Value,
	}
}

func GetLoginFromCookie(r *http.Request) string {
	//return "test@example.org"

	login, _ := r.Cookie(cookieLogin)

	logrus.Print("got login from cookies", login)

	return login.Value
}

func PutTokensToCookie(w http.ResponseWriter, tokens models.TokenPair) {
	access := http.Cookie{
		Name:  cookieAccessToken,
		Value: tokens.AccessToken,
	}
	refresh := http.Cookie{
		Name:  cookieRefreshToken,
		Value: tokens.RefreshToken,
	}

	logrus.Print("put tokens to cookies", access, refresh)

	http.SetCookie(w, &access)
	http.SetCookie(w, &refresh)
}
