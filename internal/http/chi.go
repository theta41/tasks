package http

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi/v5"
	"gitlab.com/g6834/team41/tasks/internal/env"
	"gitlab.com/g6834/team41/tasks/internal/http/letter"
	"gitlab.com/g6834/team41/tasks/internal/http/middlewares"
	"gitlab.com/g6834/team41/tasks/internal/http/task"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "gitlab.com/g6834/team41/tasks/docs"
)

func NewChi() http.Handler {
	r := chi.NewRouter()
	bindHandlers(r)

	return r
}

const (
	AcceptPath  = "/accept/{id}"
	DeclinePath = "/decline/{id}"

	CreateTaskPath = "/"
	ReadTaskPath   = "/{id}"
	UpdateTaskPath = "/{id}"
	DeleteTaskPath = "/{id}"
	ListTaskPath   = "/"
)

func bindHandlers(r *chi.Mux) {
	bindBusiness(r)
	bindProfiler(r)
	bindSwagger(r)
}

func bindBusiness(r *chi.Mux) {

	r.Route("/task", func(r chi.Router) {
		r.Use(middlewares.GetCheckAuthFunc(env.E.Auth))
		r.Post(AcceptPath, letter.Accept)
		r.Post(DeclinePath, letter.Decline)

		r.Post(CreateTaskPath, task.Create)
		r.Get(ReadTaskPath, task.Read)
		r.Put(UpdateTaskPath, task.Update)
		r.Delete(DeleteTaskPath, task.Delete)
		r.Get(ListTaskPath, task.List)
	})

}

func bindProfiler(r *chi.Mux) {
	// TODO: add profiler switch.

	r.Route("/debug/pprof", func(r chi.Router) {
		r.HandleFunc("/", pprof.Index)
		r.HandleFunc("/cmdline", pprof.Cmdline)
		r.HandleFunc("/profile", pprof.Profile)
		r.HandleFunc("/symbol", pprof.Symbol)
		r.HandleFunc("/trace", pprof.Trace)
	})
}

func bindSwagger(r *chi.Mux) {

	r.Route("/swagger", func(r chi.Router) {
		r.HandleFunc("/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
		))
	})
}
