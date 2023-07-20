package middleware

import (
	finishline "github.com/ericfisherdev/gofresh-api"
	"go-sample-app/data"
)

type Middleware struct {
	App    *finishline.FinishLine
	Models data.Models
}
