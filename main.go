package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)


func handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello, World!")
}


func main() {
	// Logger
	logger := log.New(os.Stdout, "[+]: ", log.LstdFlags)
	
	// Initialize server
	http.HandleFunc("GET /", handler)
	logger.Println("Starting server...")
	
	// Server logic
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Fatal("ListenAndServe ", err)
	}

}
