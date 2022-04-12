package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETUserId(t *testing.T) {
	t.Run("returns maru098's userViewNumber on contents 'c12'", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/contents/c12/views/maru098", nil)
		response := httptest.NewRecorder()

		UserId(response, request)

		got := response.Body.String()
		want := "3"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})
}
