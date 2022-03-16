package main

import (
	"log"
	"mini-project-go/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/", handler.HomeHandler)
	log.Println("Starting port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
