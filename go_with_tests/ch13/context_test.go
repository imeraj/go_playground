package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func assertWasCancelled(t *testing.T, s *SpyStore) {
	t.Helper()
	if s.cancelled {
		t.Errorf("store was told to cancel")
	}
}

func assertWasNotCancelled(t *testing.T, s *SpyStore) {
	t.Helper()
	if !s.cancelled {
		t.Errorf("store was not told to cancel")
	}
}

func TestHandler(t *testing.T) {
	data := "hello, world!"

	t.Run("returns data from the store", func(t *testing.T) {
		store := SpyStore{response: data}
		svr := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		assertWasCancelled(t, &store)
	})

	t.Run("tells store to cancel work if request is canceled", func(t *testing.T) {
		store := SpyStore{response: data}
		svr := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Microsecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assertWasNotCancelled(t, &store)
	})
}
