package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	//posts

	//users

	//auth

	return r

}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Handler:      mux,
		Addr:         app.config.addr,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Minute,
	}

	log.Print("Starting server on ", app.config.addr)
	return srv.ListenAndServe()
}
