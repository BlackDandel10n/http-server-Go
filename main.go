package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func createhandlerFunc(logger *log.Logger) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(writer, "Hello, World!")
		logger.Printf("Received request: %s %s from %s", req.Method, req.URL.Path, req.RemoteAddr)
	}
}


func healthcheck(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
}


func main() {
	// Logger
	logger := log.New(os.Stdout, "[+]: ", log.LstdFlags)
	
	// Initialize server
	http.HandleFunc("GET /", createhandlerFunc(logger)) // Generic pathing
	http.HandleFunc("GET /healthz", healthcheck) // Healthcheck
	logger.Println("Starting server...")
	
	// Server logic
	server := &http.Server{
		Addr: ":8080",
		Handler: nil,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	// Setup channels and OS signals
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// Shutdown logic
	go func() {
		// Listen for quit signal
		<- quit
		logger.Println("Server is shutting down...")
		
		// Shutdown context
		ctx, cancel := context.WithTimeout(context.Background(), time.Second * 30)
		defer cancel()

		server.SetKeepAlivesEnabled(false)

		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		
		close(done)
	}()

	// Run server
	logger.Printf("Server is ready to run on port %v", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %v %v\n", server.Addr, err)
	}

	<- done
	logger.Println("Server stopped")
}
