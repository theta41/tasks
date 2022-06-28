package app

import (
	"net/http"
)

type App struct {
	router http.Handler
}

func NewApp(router http.Handler) *App {
	return &App{
		router: router,
	}
}

func (a App) Run(addr string) error {
	return http.ListenAndServe(addr, a.router)
}
