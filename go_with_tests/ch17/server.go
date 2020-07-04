package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

type InMemoryPlayStore struct{}

func (i *InMemoryPlayStore) GetPlayerScore(player string) int {
	if player == "Pepper" {
		return 20
	}

	if player == "Floyd" {
		return 10
	}

	return 0
}

func (i *InMemoryPlayStore) RecordWin(name string) {

}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, score)
}
