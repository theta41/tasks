package http

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi/v5"
	"gitlab.com/g6834/team41/tasks/internal/http/letter"
	"gitlab.com/g6834/team41/tasks/internal/http/middlewares"
	"gitlab.com/g6834/team41/tasks/internal/http/task"
)

func NewChi() http.Handler {
	r := chi.NewRouter()
	bindHandlers(r)

	return r
}

const (
	AcceptPath  = "/accept/{id}"
	DeclinePath = "/decline/{id}"

	CreateTaskPath = "/task"
	ReadTaskPath   = "/task/{id}"
	UpdateTaskPath = "/task/{id}"
	DeleteTaskPath = "/task/{id}"
	ListTaskPath   = "/task"
)

func bindHandlers(r *chi.Mux) {
	bindBusiness(r)
	bindProfiler(r)
}

func bindBusiness(r *chi.Mux) {
	r.Use(middlewares.CheckAuth)

	r.Post(AcceptPath, letter.Accept)
	r.Post(DeclinePath, letter.Decline)

	r.Post(CreateTaskPath, task.Create)
	r.Get(ReadTaskPath, task.Read)
	r.Put(UpdateTaskPath, task.Update)
	r.Delete(DeleteTaskPath, task.Delete)
	r.Get(ListTaskPath, task.List)
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
