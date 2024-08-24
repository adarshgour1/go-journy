package routers

import (
	"context"
	"errors"
	"log"
	"net/http"
)

type Router struct {
	server *http.Server
	log    *log.Logger
}

func NewRouter(log *log.Logger, server *http.Server) (routers *Router) {
	routers = &Router{
		server: server,
		log:    log,
	}
	return
}

func (r *Router) Init() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
}

func (r *Router) Run() {
	go func() {
		if err := r.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			r.log.Fatalf("HTTP server error: %s", err.Error())
		}
		r.log.Println("Stopped serving new connections.")
	}()
}

func (r *Router) Close(ctx context.Context) {
	if err := r.server.Shutdown(ctx); err != nil {
		r.log.Fatalf("HTTP shutdown error: %s", err.Error())
	}
}
