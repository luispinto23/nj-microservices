package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luispinto23/go-micro/handlers"
)

// var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create handlers
	ph := handlers.NewProducts(l)

	// create mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create a new http server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start the server
	go func() {
		l.Println("starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// catch sigterm or interrupt to graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, shuting down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
