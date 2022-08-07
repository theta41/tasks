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
	AcceptPath  = "/accept/{uuid}"
	DeclinePath = "/decline/{uuid}"

	CreateTaskPath = "/"
	ListTaskPath   = "/"
	//ReadTaskPath   = "/{id}"
	//UpdateTaskPath = "/{id}"
	//DeleteTaskPath = "/{id}"
)

func bindHandlers(r *chi.Mux) {
	bindBusiness(r)
	bindProfiler(r)
	bindSwagger(r)
}

func bindBusiness(r *chi.Mux) {

	r.Route("/tasks", func(r chi.Router) {
		r.Use(middlewares.GetCheckAuthFunc(env.E.Auth))

		r.Post(CreateTaskPath, task.Create)
		r.Get(ListTaskPath, task.List)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(middlewares.TaskIdCtx)
			r.Get("/", task.Read)
			r.Put("/", task.Update)
			r.Delete("/", task.Delete)
		})

		r.Route(AcceptPath, func(r chi.Router) {
			r.Use(middlewares.TaskUuidCtx)
			r.Post("/", letter.Accept)
		})

		r.Route(DeclinePath, func(r chi.Router) {
			r.Use(middlewares.TaskUuidCtx)
			r.Post("/", letter.Decline)
		})
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
