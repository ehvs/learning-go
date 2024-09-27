package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Rebeca": 20,
			"Biles":  10,
		},
	}

	server := &PlayerServer{&store}
	t.Run("return Rebeca score", func(t *testing.T) {
		// creating instance

		// The first argument is the request's method and the second is the request's path. The nil argument refers to the request's body, which we don't need to set in this case.
		request := newGetScoreRequest("Rebeca")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponse(t, response.Body.String(), "20")
	})
	t.Run("return Biles score", func(t *testing.T) {
		// The first argument is the request's method and the second is the request's path. The nil argument refers to the request's body, which we don't need to set in this case.
		request := newGetScoreRequest("Biles")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponse(t, response.Body.String(), "10")
	})

}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func checkResponse(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
