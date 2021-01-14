package main

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app"
)

func main() {
	application, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	err = application.StartDB()
	if err != nil {
		panic(err)
	}

	err = application.Init()
	if err != nil {
		panic(err)
	}

	// Start HTTP Server
	err = application.StartHttpServer()
	if err != nil {
		panic(err)
	}
}
