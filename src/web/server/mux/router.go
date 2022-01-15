package mux

import (
	"aloksinhanov/go-bargain/src/app"
	"aloksinhanov/go-bargain/src/config"
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Server struct {
	httpServer http.Server
	Router     *mux.Router
}

var (
	server *Server
	once   sync.Once
)

func setupRouter(cfg config.ServerConfig) *mux.Router {
	mux := mux.NewRouter()
	//Add handlers here
	mux.HandleFunc(cfg.URLPrefix+cfg.APIVersion+"/ping", func(w http.ResponseWriter, r *http.Request) {
		res := "pong"
		w.Write([]byte(res))
		w.WriteHeader(http.StatusOK)
	})
	return mux
}

func NewServer(cfg config.ServerConfig) *Server {
	once.Do(func() {
		mux := setupRouter(cfg)
		server = &Server{
			Router: mux,
		}
	})
	return server
}

func (s *Server) Start(ctx context.Context, cfg config.ServerConfig) {
	wg := app.GetWaitGroup(ctx)
	wg.Add(1)
	go http.ListenAndServe(cfg.ListenURL, s.Router)
}

func (s *Server) GracefullyStop(ctx context.Context) {
	wg := app.GetWaitGroup(ctx)
	defer wg.Done()
	<-ctx.Done()

	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		log.Printf("shutdown HTTP server: %s", err)
	}
}
