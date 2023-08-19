package main

import (
	"net/http"

	"github.com/eduardylopes/bookings-app/pkg/config"
	"github.com/eduardylopes/bookings-app/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/room/majors-suite", handlers.Repo.MajorsSuite)
	mux.Get("/room/generals-quarters", handlers.Repo.GeneralsQuarters)
	mux.Get("/room/reservation", handlers.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
