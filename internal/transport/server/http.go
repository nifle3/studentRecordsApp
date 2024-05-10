package server

import (
	"net/http"
	"time"
)

type Opts struct {
	Port        string
	IdleTimeout time.Duration
}

type Server struct {
	Opts
}

func New(opts Opts) Server {
	return Server{
		Opts: opts,
	}
}

func (s Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/auth")

	mux.HandleFunc("GET /v1/admin")
	mux.HandleFunc("PATCH /v1/admin")
	mux.HandleFunc("GET /v1/admin/worker")
	mux.HandleFunc("GET /v1/admin/worker/{id}")
	mux.HandleFunc("POST /v1/admin/worker")
	mux.HandleFunc("PATCH /v1/admin/worker")

	mux.HandleFunc("GET /v1/worker")
	mux.HandleFunc("GET /v1/worker/student")
	mux.HandleFunc("GET /v1/worker/student/{id}")
	mux.HandleFunc("POST /v1//worker/student")
	mux.HandleFunc("PATCH /v1/worker/student")
	mux.HandleFunc("GET /v1/worker/application")
	mux.HandleFunc("GET /v1/worker/application/{userId}")
	mux.HandleFunc("GET /v1/worker/application/{id}")
	mux.HandleFunc("PATCH /v1/worker/application/close")
	mux.HandleFunc("GET /v1/worker/document/download/{userId}")
	mux.HandleFunc("GET /v1/worker/document/{userId}")
	mux.HandleFunc("GET /v1/worker/document/{Id}")
	mux.HandleFunc("POST /v1/worker/document")
	mux.HandleFunc("PATCH /v1/worker/document}")

	mux.HandleFunc("GET /v1/student")
	mux.HandleFunc("GET /v1/student/docuemnt")
	mux.HandleFunc("GET /v1/student/application")
	mux.HandleFunc("PATCH /v1/student/application")
	mux.HandleFunc("POST /v1/student/application")

	server := http.Server{
		Addr:        s.Port,
		IdleTimeout: s.IdleTimeout,
		Handler:     mux,
	}

	return server.ListenAndServe()
}
