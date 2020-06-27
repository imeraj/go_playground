package main

import ( 
	"time"
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestRacer(t *testing.T) {
	t.Run("no timeout", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
	
		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("timeout", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)
	
		defer server.Close()
	
		_, err := ConfigurableRacer(server.URL, server.URL,  20 * time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
