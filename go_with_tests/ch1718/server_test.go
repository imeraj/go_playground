package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
	league []Player
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	score := s.scores[player]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.scores[name]++
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
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

	server := NewPlayerServer(&store)

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
		nil,
	}

	server := NewPlayerServer(&store)

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

func TestRecodringWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)
	server.ServeHTTP(response, request)
	server.ServeHTTP(response, request)

	t.Run("get score", func(t *testing.T) {
		request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
		response = httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "3"

		assertStatus(t, response.Code, http.StatusOK)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("get league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player
		want := []Player{
			{"Pepper", 3},
		}

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}

	store := StubPlayerStore{nil, wantedLeague}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)

		if response.Result().Header.Get("Content-Type") != "application/json" {
			t.Errorf("resonse did not have content-type of application/json, got %v", response.Result().Header)
		}

		if !reflect.DeepEqual(got, wantedLeague) {
			t.Errorf("got %v, want %v", got, wantedLeague)
		}
	})
}
