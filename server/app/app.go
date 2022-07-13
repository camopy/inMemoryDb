package app

import "github.com/camopy/mosaico/adapter/db"

type App struct {
	db db.DB
}

func New(db db.DB) *App {
	return &App{
		db: db,
	}
}
