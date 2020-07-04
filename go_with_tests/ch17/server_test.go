package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	score := s.scores[player]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.scores[name]++
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertStatus(t, response.Code, http.StatusOK)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertStatus(t, response.Code, http.StatusOK)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns 404 on missing player", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Apollo", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestStoreWins(t *testing.T) {
	player := "Pepper"

	store := StubPlayerStore{
		map[string]int{},
	}
	server := &PlayerServer{&store}

	t.Run("it records wins when  POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if store.scores[player] != 1 {
			t.Errorf("got %d, want %d", store.scores[player], 1)
		}
	})
}
