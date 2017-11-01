package main

import (
	"log"
	"net/http"

	"github.com/acoshift/middleware"
)

func main() {
	m := http.NewServeMux()
	m.Handle("/", http.FileServer(web{"dist"}))
	m.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	h := middleware.Chain(
		cacheControl,
	)(m)

	log.Println("Start Web Sever on :8080")
	log.Fatal(http.ListenAndServe(":8080", h))
}

func cacheControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=3600")
		h.ServeHTTP(w, r)
	})
}

type web struct {
	dir http.Dir
}

func (w web) Open(name string) (http.File, error) {
	fs, err := w.dir.Open(name)
	if err != nil {
		return nil, err
	}
	return fs, nil
}
