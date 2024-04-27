package server

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"studentRecordsApp/internal/service"
)

var jwtSecretKey = []byte("very-secret-key")

type jwtClaims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
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
}

func New(application service.Application, student service.Student, phoneNumber service.PhoneNumber,
	document service.Document, user service.User) *Server {

	return &Server{
		application: application,
		student:     student,
		phoneNumber: phoneNumber,
		document:    document,
		user:        user,
	}
}

func (s *Server) Start() error {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	v1Api := r.Group("/api/v1")
	{
		v1Api.POST("/login", s.auth)

		studentGroup := v1Api.Group("/student")
		{
			studentGroup.Use(s.authMiddleware(roleStudent))
			studentGroup.GET("/self", s.GetSelfStudent)
			studentGroup.PUT("/student", s.UpdateSelfStudent)
			studentGroup.GET("/applications", s.GetAllApplications)
			studentGroup.GET("/application/:id", s.GetApplication)
			studentGroup.POST("/application", s.CreateApplication)
			studentGroup.PUT("/application/:id", s.UpdateApplication)
			studentGroup.GET("/document", s.AllDocumentsForStudent)
			studentGroup.GET("/document/:id", s.GetDocument)
		}

		adminGroup := v1Api.Group("/admin")
		{
			adminGroup.Use(s.authMiddleware(roleAdmin))
			adminGroup.GET("/self", s.GetUserSelfAccount)
			adminGroup.GET("/worker", s.GetAllWorker)
			adminGroup.GET("/worker/:id", s.GetWorkerById)
			adminGroup.POST("/worker", s.AddWorker)
			adminGroup.POST("/worker/:id", s.UpdateWorker)
			adminGroup.DELETE("/worker/:id", s.DeleteWorker)
		}

		workerGroup := v1Api.Group("/worker")
		{
			workerGroup.Use(s.authMiddleware(roleWorker))
			workerGroup.GET("/self", s.GetUserSelfAccount)
			workerGroup.GET("/student", s.GetAllStudent)
			workerGroup.GET("/student/:id", s.GetStudentById)
			workerGroup.POST("/student", s.AddStudent)
			workerGroup.POST("/student/:id", s.UpdateStudent)
			workerGroup.DELETE("/student/:id", s.DeleteStudent)
			workerGroup.POST("/student/:id/phone", s.AddPhoneNumber)
			workerGroup.POST("/student/:id/document", s.AddDocument)
			workerGroup.PATCH("/student/phone", s.UpdatePhoneNumber)
			workerGroup.DELETE("/student/:id/phone/:phoneId", s.DeletePhoneNumber)
			workerGroup.DELETE("/student/:id/document/:documentId", s.DeleteDocument)
			workerGroup.GET("/student/:id/phone", s.GetPhoneNumbers)
			workerGroup.GET("/student/:id/document", s.GetDocuments)
			workerGroup.GET("/student/:id/document/:documentId", s.GetDocumentById)
			workerGroup.GET("/student/:id/application", s.GetApplications)
			workerGroup.GET("/student/:id/application/:applicationId", s.GetApplicationById)
			workerGroup.PATCH("/student/:id/application/:applicationId", s.CloseApplication)
		}
	}

	return r.Run()
}
