package handlers

import (
	finishline "github.com/ericfisherdev/gofresh-api"
	"go-sample-app/data"
	"net/http"
	"time"
)

type Handlers struct {
	App    *finishline.FinishLine
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering:", err)
	}
}
