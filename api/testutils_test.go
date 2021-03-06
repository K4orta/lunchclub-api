package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/k4orta/lunchclub-api/storage"
)

func init() {
	storage.SetupDBForTesting()
}

func makeFakeServer(resp string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, resp)
	}))
}

func mockAPI(handler func(http.ResponseWriter, *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(handler))
}
