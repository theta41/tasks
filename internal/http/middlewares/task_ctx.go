package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"gitlab.com/g6834/team41/tasks/internal/http/util"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

const ContextKeyTaskId = util.ContextKey("id")
const ContextKeyTaskUuid = util.ContextKey("uuid")

func TaskIdCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log := logrus.WithField("middleware", "TaskIdCtx")

		taskId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			log.Error("missing task id: ", err)
			return
		}

		//log.Infof("got task id %v", taskId)
		ctx := context.WithValue(r.Context(), ContextKeyTaskId, taskId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TaskUuidCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log := logrus.WithField("middleware", "TaskUuidCtx")

		uuid := chi.URLParam(r, "uuid")
		if uuid == "" {
			http.Error(w, "", http.StatusBadRequest)
			log.Error("missing letter uuid")
			return
		}

		//log.Infof("got task uuid %v", uuid)
		ctx := context.WithValue(r.Context(), ContextKeyTaskUuid, uuid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
