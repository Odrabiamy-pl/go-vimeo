package mocks

import (
	_ "embed"
	"fmt"
	"net/http"
	"net/http/httptest"
)

//go:embed videos-899476146.json
var getVideos899476146 []byte

//go:embed videos.json
var getVideos []byte

func Run() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		methodWithUri := fmt.Sprintf("%s %s", r.Method, r.RequestURI)
		switch methodWithUri {
		case "GET /users/202422358/videos":
			w.Header().Add("Content-Type", "application/vnd.vimeo.video+json")
			if _, err := w.Write(getVideos); err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
		case "GET /videos/899476146":
			w.Header().Add("Content-Type", "application/vnd.vimeo.video+json")
			if _, err := w.Write(getVideos899476146); err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
}
