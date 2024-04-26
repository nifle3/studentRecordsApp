package server

import (
	"embed"
	"net/http"
	"studentRecordsApp/internal/service"
)

type Server struct {
	application service.Application
	student     service.Student
	phoneNumber service.PhoneNumber
	document    service.Document
	user        service.User
	fs          *embed.FS
}

func New(application service.Application, student service.Student, phoneNumber service.PhoneNumber,
	document service.Document, user service.User, fs *embed.FS) *Server {

	return &Server{
		application: application,
		student:     student,
		phoneNumber: phoneNumber,
		document:    document,
		user:        user,
		fs:          fs,
	}
}

func (s *Server) Start() error {

	return http.ListenAndServe(":8080", nil)
}
