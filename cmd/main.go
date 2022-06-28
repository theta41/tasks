package main

import (
	"main/internal/app"
	"main/internal/env"
	"main/internal/http"
)

func main() {
	r := http.NewChi()
	a := app.NewApp(r)
	err := a.Run(env.E.C.HostAddress)
	if err != nil {
		panic(err)
	}
}
