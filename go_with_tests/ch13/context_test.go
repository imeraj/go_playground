package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestHandler(t *testing.T) {
	data := "hello, world"
	svr := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	resp_writer := httptest.NewRecorder()

	svr.ServeHTTP(resp_writer, request)

	if resp_writer.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, resp_writer.Body.String(), data)
	}
}
