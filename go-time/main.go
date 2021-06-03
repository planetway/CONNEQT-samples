package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	t := time.Now()
	enc := json.NewEncoder(w)
	err := enc.Encode(map[string]interface{}{
		"now": t,
	})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
