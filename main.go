package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const ApiUrl = "/api/games/"

type GameModel struct {
	ID        int
	State     string
	WonByTeam string `json:",omitempty"`
}

var games []*Game

func BadRequest(w http.ResponseWriter) {
	http.Error(w, fmt.Sprintf("%d %s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)), http.StatusBadRequest)
}

func toModel(g *Game) *GameModel {
	model := GameModel{
		ID:    g.id,
		State: fmt.Sprint(g.State()),
	}
	if g.State() == Finished {
		if g.WonByTeam(TeamA) {
			model.WonByTeam = fmt.Sprint(TeamA)
		} else {
			model.WonByTeam = fmt.Sprint(TeamB)
		}
	}
	return &model
}

func findGame(id int) *Game {
	for _, g := range games {
		if g.id == id {
			return g
		}
	}
	return nil
}

func handleGame(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)
	e.SetIndent("", "\t")

	log.Printf("Handling request: %s %s", r.Method, r.RequestURI)

	log.Printf("LOCATION: %s", r.URL)

	switch r.Method {
	case http.MethodGet:
		param := strings.Trim(strings.TrimPrefix(r.RequestURI, ApiUrl), "/")
		var ID int
		if param != "" && len(param) > 0 {
			id, err := strconv.Atoi(param)
			if err != nil {
				BadRequest(w)
				return
			}
			ID = id
			log.Printf("ID: %d", ID)
		}
		if ID > 0 {
			var game = findGame(ID)
			if game != nil {
				e.Encode(toModel(game))
			} else {
				http.NotFound(w, r)
			}
		} else {
			var gamesModel = make([]*GameModel, 0)
			for _, g := range games {
				gamesModel = append(gamesModel, toModel(g))
			}
			e.Encode(gamesModel)
		}
	case http.MethodPost:
		path := strings.Trim(strings.TrimPrefix(r.RequestURI, ApiUrl), "/")
		params := strings.Split(path, "/goal/")
		var game *Game
		if path != "" {
			var ID int
			var team Team
			if len(params) < 2 {
				BadRequest(w)
				return
			}
			id, err := strconv.Atoi(params[0])
			if err != nil {
				BadRequest(w)
				return
			}
			ID = id
			if params[1] == "TeamA" {
				team = TeamA
			} else if params[1] == "TeamB" {
				team = TeamB
			} else {
				BadRequest(w)
				return
			}
			game = findGame(ID)
			if game == nil {
				http.NotFound(w, r)
				return
			}
			game.AddScore(team)
		} else {
			game = NewGame()
			games = append(games, game)
			w.Header().Set("Location", fmt.Sprintf("%s%d", ApiUrl, game.id))
			w.WriteHeader(http.StatusCreated)
		}
		e.Encode(toModel(game))
	default:
		http.NotFound(w, r)
	}
}

func main() {
	const port = 8080

	mux := http.NewServeMux()
	mux.HandleFunc(ApiUrl, handleGame)

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
