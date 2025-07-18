package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)


func CreatehandlerFunc(logger *log.Logger) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(writer, "Hello, World!")
		logger.Printf("Received request: %s %s from %s", req.Method, req.URL.Path, req.RemoteAddr)
	}
}


func main() {
	// Logger
	logger := log.New(os.Stdout, "[+]: ", log.LstdFlags)
	
	// Initialize server
	http.HandleFunc("GET /", CreatehandlerFunc(logger))
	logger.Println("Starting server...")
	
	// Server logic
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Fatal("ListenAndServe ", err)
	}

}
