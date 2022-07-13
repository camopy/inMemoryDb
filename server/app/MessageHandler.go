package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Message struct {
	Id      uint64 `json:"id"`
	Message string `json:"message"`
}

type notFound struct {
	Err string `json:"err"`
}

func (a *App) HandleGetMessage(w http.ResponseWriter, r *http.Request) {
	id, err := parseIdFromURL(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg, err := a.db.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, &notFound{Err: "not found"})
		return
	}

	m := &Message{
		Id:      id,
		Message: msg,
	}

	render.JSON(w, r, m)
}

func (app *App) HandleNewMessage(w http.ResponseWriter, r *http.Request) {
	id, err := parseIdFromURL(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message := &Message{
		Id: id,
	}

	if err := json.NewDecoder(r.Body).Decode(message); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	app.db.Set(id, message.Message)
	render.JSON(w, r, message)
}

func parseIdFromURL(r *http.Request) (uint64, error) {
	return strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
}
