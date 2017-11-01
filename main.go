package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	srv := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}

func cacheControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=7200")
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
