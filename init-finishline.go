package main

import (
	finishline "github.com/ericfisherdev/gofresh-api"
	"go-sample-app/data"
	"go-sample-app/handlers"
	"go-sample-app/middleware"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init finishline
	fin := &finishline.FinishLine{}
	err = fin.New(path)
	if err != nil {
		log.Fatal(err)
	}

	fin.AppName = "FinishLine"

	myMiddleware := &middleware.Middleware{
		App: fin,
	}

	myHandlers := &handlers.Handlers{
		App: fin,
	}

	app := &application{
		App:        fin,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)

	myHandlers.Models = app.Models

	app.Middleware.Models = app.Models

	return app
}
