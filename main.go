package main

import (
	"context"
	"fxproxy/pkg/apis"
	"fxproxy/pkg/logger"
	"fxproxy/pkg/validator"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger.Start()
	http.HandleFunc("/", apis.CoreHandler)
	logger.LogInfo("started listening on port " + os.Getenv("PORT"))
	logger.LogInfo("Waiting request...")
	validator.GlobalMatcher = validator.NewMatcher(validator.AllowedList)
	httpServer := &http.Server{
		Addr: os.Getenv("PORT"),
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	ctx, signalCancel := signal.NotifyContext(context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	defer signalCancel()

	<-ctx.Done()

	log.Print("os.Interrupt - shutting down...\n")

	Context, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(Context); err != nil {
		log.Printf("shutdown error: %v\n", err)
		os.Exit(1)
	}
	log.Printf("Context stopped\n")
}
