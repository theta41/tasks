package main

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/app"
	"gitlab.com/g6834/team41/tasks/internal/env"
	"gitlab.com/g6834/team41/tasks/internal/http"
)

func main() {
	logrus.Info("Starting server...")
	r := http.NewChi()

	a := app.NewApp(r)
	err := a.Run(env.E.C.HostAddress)
	if err != nil {
		logrus.Panic(err)
	}
}
