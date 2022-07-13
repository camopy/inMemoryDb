package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/camopy/mosaico/adapter/db"
	"github.com/camopy/mosaico/server/app"
	"github.com/camopy/mosaico/server/router"
)

func main() {
	app := initApp()
	srv := initServer(app)
	startServer(app, srv)
}

func initApp() *app.App {
	return app.New(db.New())
}

func initServer(app *app.App) *http.Server {
	r := router.New(app)

	return &http.Server{
		Addr:    ":8085",
		Handler: r,
	}
}

func startServer(app *app.App, srv *http.Server) {
	fmt.Println("Starting server")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
