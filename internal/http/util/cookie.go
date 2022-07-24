package util

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"gitlab.com/g6834/team41/tasks/internal/models"
)

const (
	cookieLogin        = "Login"
	cookieAccessToken  = "AccessToken"
	cookieRefreshToken = "RefreshToken"
)

func mustCookie(r *http.Request, name string) string {
	v, err := r.Cookie(name)
	if err != nil || v == nil {
		logrus.Infof("missing cookie %s", name)
		return ""
	}

	// logrus.Debugf ?
	logrus.Infof("got cookie %v", v)
	return v.Value
}

func GetTokensFromCookie(r *http.Request) models.TokenPair {

	access := mustCookie(r, cookieAccessToken)
	refresh := mustCookie(r, cookieRefreshToken)

	return models.TokenPair{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func GetLoginFromCookie(r *http.Request) string {
	//return "test@example.org"

	login := mustCookie(r, cookieLogin)

	return login
}

func PutTokensToCookie(w http.ResponseWriter, tokens models.TokenPair) {
	access := http.Cookie{
		Name:    cookieAccessToken,
		Value:   tokens.AccessToken,
		Path:    "/",
		Expires: time.Time{}.AddDate(9998, 0, 0), //learning cookies never expires
	}
	refresh := http.Cookie{
		Name:    cookieRefreshToken,
		Value:   tokens.RefreshToken,
		Path:    "/",
		Expires: time.Time{}.AddDate(9998, 0, 0), //learning cookies never expires
	}

	//logrus.Debugf("put tokens to cookie %+v %+v", access, refresh)
	logrus.Infof("put tokens to cookie %v %v", access, refresh)

	http.SetCookie(w, &access)
	http.SetCookie(w, &refresh)
}

func PutLoginToCookie(w http.ResponseWriter, loginValue string) {
	loginCookie := http.Cookie{
		Name:    cookieLogin,
		Value:   loginValue,
		Path:    "/",
		Expires: time.Time{}.AddDate(9998, 0, 0), //learning cookies never expires
	}

	logrus.Infof("put login to cookie %v", loginCookie)

	http.SetCookie(w, &loginCookie)
}
