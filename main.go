package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Privet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Zapusk web-server")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

	fmt.Println("Hello, World!")
}
