package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/salamwaddah/full-courses/go/building-microservices/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", ph)

	// create a new http server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// wrap ListenAndServe in a goroutine, so it doesn't block the graceful shutdown
	go func() {
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
