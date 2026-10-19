package main

import (
	"log"
	"net/http"
	"encoding/json"
)


type Sentence struct {
	Message string `json:"text"`
}

type Vec struct {
	Numbers []float32 `json:"embedding"`
}

//handleEmbed receives an embed request and replies with the corresponding vector if valid
func handleEmbed(w http.ResponseWriter, r *http.Request){

	//checks if HTTP request type is POST
	if r.Method != http.MethodPost {
		http.Error(w, "not a valid request", http.StatusMethodNotAllowed)
		return
	}

	var req Sentence

	//read the JSON from the request and fill in req
	err := json.NewDecoder(r.Body).Decode(&req)

	//handles bad json and sends 400sc if so
	if err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	log.Println(req.Message)

	//example vector encoding as JSON and written to response
	var resp Vec
	resp.Numbers = []float32{0.1, 0.2, 0.3}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	

}

func main() {
	//route requests for /embed to handleEmbed
	http.HandleFunc("/embed", handleEmbed)
	log.Fatal(http.ListenAndServe(":8080", nil))
}