package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/salamwaddah/full-courses/go/building-microservices/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProducts(l)

	// create a serve mux
	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareValidateProduct)

	// create a new http server
	port := ":9090"
	s := &http.Server{
		Addr:         port,
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// wrap ListenAndServe in a goroutine, so it doesn't block the graceful shutdown
	go func() {
		l.Println("Starting server on port", port)

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// create a signal channel to listen for os interruptions
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// block until a signal is received
	sig := <-sigChan

	// setup graceful shutdown
	l.Println("Received terminate, graceful shutdown", sig)
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeoutContext)
}
