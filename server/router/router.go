package router

import (
	"github.com/camopy/mosaico/server/app"
	"github.com/go-chi/chi"
)

func New(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/{id}", app.HandleGetMessage)
	r.Post("/{id}", app.HandleNewMessage)

	return r
}
