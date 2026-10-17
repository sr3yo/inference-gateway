package main

import (
	"log"
	"net/http"
)

func handleEmbed(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/embed", handleEmbed)
	log.Fatal(http.ListenAndServe(":8080", nil))
}