package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
	"github.com/sarulabs/di-example/app/handlers"
	"github.com/sarulabs/di-example/app/middlewares"
	"github.com/sarulabs/di-example/config/logging"
	"github.com/sarulabs/di-example/config/services"
)

func main() {
	// Use a single logger in the whole application.
	// Need to close it at the end.
	defer logging.Logger.Sync()

	// Create the app container.
	// Do not forget to delete it at the end.
	builder, err := di.NewBuilder()
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	err = builder.Add(services.Services...)
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	app := builder.Build()
	defer app.Delete()

	// Create the http server.
	r := mux.NewRouter()

	// Function to apply the middlewares:
	// - recover from panic
	// - add the container in the http requests
	m := func(h http.HandlerFunc) http.HandlerFunc {
		return middlewares.PanicRecoveryMiddleware(
			di.HTTPMiddleware(h, app, func(msg string) {
				logging.Logger.Error(msg)
			}),
			logging.Logger,
		)
	}

	r.HandleFunc("/cars", m(handlers.GetCarListHandler)).Methods("GET")
	r.HandleFunc("/cars", m(handlers.PostCarHandler)).Methods("POST")
	r.HandleFunc("/cars/{carId}", m(handlers.GetCarHandler)).Methods("GET")
	r.HandleFunc("/cars/{carId}", m(handlers.PutCarHandler)).Methods("PUT")
	r.HandleFunc("/cars/{carId}", m(handlers.DeleteCarHandler)).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + os.Getenv("SERVER_PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logging.Logger.Info("Listening on port " + os.Getenv("SERVER_PORT"))

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Logger.Error(err.Error())
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	logging.Logger.Info("Stopping the http server")

	if err := srv.Shutdown(ctx); err != nil {
		logging.Logger.Error(err.Error())
	}
}
