package main

import (
	"log"
	"net/http"
)


type Sentence struct {
	Message string `json:"text"`
}

//handleEmbed receives an embed request and replies with the corresponding vector if valid
func handleEmbed(w http.ResponseWriter, r *http.Request){

	//checks if HTTP request type is POST
	if r.Method != http.MethodPost {
		http.Error(w, "not a valid request", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	//route requests for /embed to handleEmbed
	http.HandleFunc("/embed", handleEmbed)
	log.Fatal(http.ListenAndServe(":8080", nil))
}