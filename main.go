package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GameModel struct {
}

func handleGame(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)
	log.Printf("Handling request: %s %s", r.Method, r.RequestURI)
	switch r.Method {
	case http.MethodGet:
		e.Encode(&GameModel{})
		break
	case http.MethodPost:
	default:
		http.Error(w, "404 Not found", http.StatusNotFound)
	}
}

func main() {
	const port = 8080

	mux := http.NewServeMux()
	mux.HandleFunc("/api/games/", handleGame)

	log.Printf("Server starting on port %v\n", port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
