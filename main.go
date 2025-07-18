package main

import (
	"fmt"
	"net/http"
)


func handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello, World!")
}


func main() {
	http.HandleFunc("GET /", handler)
	http.ListenAndServe(":8080", nil)
}
