package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/env"
	"gitlab.com/g6834/team41/tasks/internal/http/util"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO: Get login from cookie
		login := "test@example.org"

		tokens := util.GetTokensFromCookie(r)

		// Validate access rights
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		newTokens, err := env.E.Auth.Validate(ctx, login, tokens)
		if err != nil {
			http.Error(w, "{}", http.StatusForbidden)
			logrus.Error(err)
			return
		}

		util.PutTokensToCookie(w, newTokens)
		next.ServeHTTP(w, r)
	})
}
