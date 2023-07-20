package main

import (
	finishline "github.com/ericfisherdev/gofresh-api"
	"go-sample-app/data"
	"go-sample-app/handlers"
	"go-sample-app/middleware"
)

type application struct {
	App        *finishline.FinishLine
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	f := initApplication()
	f.App.ListenAndServe()
}
