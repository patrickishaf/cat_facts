package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/patrickishaf/cat_facts/lib"
)

type apiResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type response struct {
	Status string `json:"status"`
	User   struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Stack string `json:"stack"`
	} `json:"user"`
	Timestamp time.Time `json:"timestamp"`
	Fact      string    `json:"fact"`
}

func sendResponse(w *http.ResponseWriter, fact string) {
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(response{
		Status: "success",
		User: struct {
			Email string "json:\"email\""
			Name  string "json:\"name\""
			Stack string "json:\"stack\""
		}{
			Email: "onumdev@gmail.com",
			Name:  "Peter Ndukwe",
			Stack: "Go/stdlib",
		},
		Timestamp: time.Now(),
		Fact:      fact,
	})
}

func fetchCatFact() string {
	res, err := lib.NewHttpClient().Get("https://catfact.ninja/fact")
	if err != nil {
		return "the api is currently unavailable"
	}
	var fact apiResponse
	if err := json.Unmarshal(*res, &fact); err != nil {
		return "something went wrong"
	}
	return fact.Fact
}

func main() {
	http.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		data := fetchCatFact()
		sendResponse(&w, data)
	})
	http.ListenAndServe(":8080", nil)
}
