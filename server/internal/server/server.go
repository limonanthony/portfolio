package server

import (
	"net/http"
	"time"

	"github.com/limonanthony/portfolio/internal/router"
)

const serverReadTimeout = 10 * time.Second
const serverWriteTimeout = 10 * time.Second
const serverIdleTimeout = 60 * time.Second

type Server struct {
	Router *router.Router
	http   *http.Server
}

func NewServer() Server {
	r := router.NewRouter()

	return Server{
		Router: r,
		http: &http.Server{
			Addr:         ":8080",
			Handler:      r,
			ReadTimeout:  serverReadTimeout,
			WriteTimeout: serverWriteTimeout,
			IdleTimeout:  serverIdleTimeout,
		},
	}
}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}
