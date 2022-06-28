package http

import (
	"github.com/go-chi/chi/v5"
	"main/internal/http/letter"
	"main/internal/http/task"
	"net/http"
	"net/http/pprof"
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
