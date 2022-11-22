package main

import (
	"context"
	"fmt"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func ListenAndServe(addr string, handler http.Handler) error {
	log.Infof("Listening on Port: %v", addr)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", addr),
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	wgServer := sync.WaitGroup{}
	wgServer.Add(2)

	var serverError error

	go func() {
		defer wgServer.Done()
		killSignal := <-signals
		switch killSignal {
		case os.Interrupt:
			log.Infoln("SIGINT recieved (Control-C ?)")
		case syscall.SIGTERM:
			log.Infoln("SIGTERM recieved (Heroku shutdown?)")
		case nil:
			return
		}
		log.Infoln("graceful shutdown...")
		err := srv.Shutdown(context.Background())
		if err != nil {
			log.Error(err.Error())
		}
		log.Infoln("graceful shutdown complete")
	}()

	go func() {
		defer wgServer.Done()
		if err := srv.ListenAndServe(); err != nil {
			serverError = err
		}
		signals <- nil
	}()

	wgServer.Wait()
	return serverError
}

func CorsHandler() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "X-Requested-With", "Authorization", "Content-Type", "X-Requested-With", "Bearer", "Origin"},
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})
}
