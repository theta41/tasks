package main

import (
	"gitlab.com/g6834/team41/tasks/internal/app"
	"gitlab.com/g6834/team41/tasks/internal/env"
	"gitlab.com/g6834/team41/tasks/internal/http"
)

func main() {
	r := http.NewChi()
	a := app.NewApp(r)
	err := a.Run(env.E.C.HostAddress)
	if err != nil {
		panic(err)
	}
}
