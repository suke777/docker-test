package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("start server.")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("received request.")
	fmt.Fprintf(w, "Hello World")
}
