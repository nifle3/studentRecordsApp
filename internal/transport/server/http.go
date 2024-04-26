package server

import (
	"embed"
	"github.com/golang-jwt/jwt"
	"net/http"

	"studentRecordsApp/internal/service"
)

var jwtSecretKey = []byte("very-secret-key")

type jwtClaims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.Claims
}

const (
	roleWorker  = "employee"
	roleAdmin   = "admin"
	roleStudent = "student"

	tokenCookie = "token"
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

	http.HandleFunc("POST /v1/auth", s.auth)

	http.HandleFunc("GET /", s.authPage)

	return http.ListenAndServe(":8080", nil)
}
