package server

import (
	"net/http"
	"studentRecordsApp/internal/service"
)

type Server struct {
    application service.Application
    student     service.Student
    phoneNumber service.PhoneNumber
    document    service.Document
    user        service.User
}

func (s *Server) Start() error {

    return http.ListenAndServe(":8080", nil)
}
