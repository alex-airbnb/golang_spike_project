package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alex-airbnb/golang_spike_project/router"
)

func main() {
	muxRouter := router.GetRouter()

	server := http.Server{
		Addr:         ":8000",
		Handler:      muxRouter,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Println("golang_spike_project is running on port 8000")

		err := server.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}()

	shutdownChannel := make(chan os.Signal)

	signal.Notify(shutdownChannel, os.Interrupt)
	signal.Notify(shutdownChannel, os.Kill)

	shutdownNotification := <-shutdownChannel

	log.Println("Received shutdown notification", shutdownNotification)

	timeout, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	server.Shutdown(timeout)
}
