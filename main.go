package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Sentence struct {
	Message string `json:"text"`
}

type Vec struct {
	Numbers []float32 `json:"embedding"`
}

type Batch struct {
	Words []string `json:"texts"`
}

type Embeds struct {
	Vecs [][]float32 `json:"embeddings"`
}

// handleEmbed receives an embed request and replies with the corresponding vector if valid
func handleEmbed(w http.ResponseWriter, r *http.Request) {

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

	//handles empty message
	if req.Message == "" {
		http.Error(w, "Empty message", http.StatusBadRequest)
		return
	}

	var pyReq Batch
	pyReq.Words = []string{req.Message}

	body, err := json.Marshal(pyReq)
	if err != nil {
		http.Error(w, "could not encode worker request",
			http.StatusInternalServerError)
		return
	}

	//creates a POST request containing the JSON body for the Python worker
	workerReq, err := http.NewRequestWithContext(
		r.Context(),
		http.MethodPost,
		"http://127.0.0.1:8000/embed",
		bytes.NewReader(body),
	)
	if err != nil {
		http.Error(w, "could not create worker request",
			http.StatusInternalServerError)
		return
	}

	//instantiates an HTTP client and sends the POST request to the Python worker
	workerReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}

	workerResp, err := client.Do(workerReq)
	if err != nil {
		http.Error(w, "worker unavailable", http.StatusBadGateway)
		return
	}
	defer workerResp.Body.Close()

	//checks that Python returned a successful HTTP status
	if workerResp.StatusCode != http.StatusOK {
		http.Error(w, "worker failed", http.StatusBadGateway)
		return
	}

	//decodes Python's JSON into a Go Struct
	var result Embeds
	if err := json.NewDecoder(workerResp.Body).Decode(&result); err != nil {
		http.Error(w, "invalid worker response", http.StatusBadGateway)
		return
	}

	//checks that Python returned one vector containing 384 numbers
	if len(result.Vecs) != 1 || len(result.Vecs[0]) != 384 {
		http.Error(w, "unexpected embedding dimensions", http.StatusBadGateway)
		return
	}

	//return the real vector from Python
	resp := Vec{Numbers: result.Vecs[0]}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func main() {
	//route requests for /embed to handleEmbed
	http.HandleFunc("/embed", handleEmbed)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
