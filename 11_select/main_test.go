package sselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("fastest server wins", func(t *testing.T) {
		slowServer := makeDelayedServer(20)
		defer slowServer.Close()
		fastServer := makeDelayedServer(0)
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(fastUrl, slowUrl)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("error if no response in 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(11_000)
		defer slowServer.Close()
		slowerServer := makeDelayedServer(12_000)
		defer slowerServer.Close()

		_, err := configurableRacer(slowServer.URL, slowerServer.URL, 10*time.Millisecond)

		if nil == err {
			t.Error("expected an error")
		}
	})
}

func makeDelayedServer(delay int) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	return server
}
