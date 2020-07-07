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
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r, player)
	case http.MethodGet:
		p.showScore(w, r, player)
	}
}

type InMemoryPlayStore struct {
	store map[string]int
}

func NewInMemoryPlayStore() *InMemoryPlayStore {
	return &InMemoryPlayStore{store: map[string]int{}}
}

func (i *InMemoryPlayStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayStore) RecordWin(name string) {
	i.store[name]++
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, score)
}
