package server

import (
	"fmt"
	"net/http"
	entities "studentRecordsApp/internal/service/entities"
	"time"

	"studentRecordsApp/internal/service"
)

type Opts struct {
	Port        string
	IdleTimeout time.Duration
}

func NewOtps(port string, IdleTimeout time.Duration) Opts {
	return Opts{
		Port:        port,
		IdleTimeout: IdleTimeout,
	}
}

type Mux struct {
	application service.Application
	document    service.Document
	auth        service.Auth
	student     service.Student
	user        service.User
}

func NewMux(application service.Application, document service.Document, auth service.Auth,
	student service.Student, user service.User) Mux {
	return Mux{
		application: application,
		document:    document,
		auth:        auth,
		student:     student,
		user:        user,
	}
}

type Server struct {
	Mux
	Opts
}

func New(opts Opts, mux Mux) Server {
	return Server{
		Mux:  mux,
		Opts: opts,
	}
}

func (s Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/auth", s.Login)
	mux.HandleFunc("GET /v1/role", s.GetRole)
	mux.HandleFunc("GET /v1/role/worker", s.GetWorkerRole)

	mux.HandleFunc("GET /v1/admin", s.SecureHandler(entities.UserAdmin, s.AdminGetSelf))
	mux.HandleFunc("PATCH /v1/admin", s.SecureHandler(entities.UserAdmin, s.AdminPatchSelf))
	mux.HandleFunc("GET /v1/admin/worker", s.SecureHandlerWithOutId(entities.UserAdmin, s.AdminGetAllWorkers))
	mux.HandleFunc("POST /v1/admin/worker", s.SecureHandlerWithOutId(entities.UserAdmin, s.AdminAddWorker))
	mux.HandleFunc("PATCH /v1/admin/worker/{id}", s.SecureHandlerWithOutId(entities.UserAdmin, s.AdminPatchWorker))
	mux.HandleFunc("DELETE /v1/admin/worker/{id}", s.SecureHandlerWithOutId(entities.UserAdmin, s.AdminDeleteWorker))

	mux.HandleFunc("GET /v1/worker", s.SecureHandler(entities.UserWorker, s.WorkerGetSelf))
	mux.HandleFunc("GET /v1/worker/student", s.SecureHandlerWithOutId(entities.UserWorker, s.WorkerGetAllStudents))
	mux.HandleFunc("GET /v1/worker/student/{id}", s.SecureHandlerWithOutId(entities.UserWorker, s.WorkerGetStudent))
	mux.HandleFunc("GET /v1/worker/student/image/{link}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerGetStudentImage))
	mux.HandleFunc("POST /v1/worker/student", s.SecureHandlerWithOutId(entities.UserWorker, s.WorkerAddStudent))
	mux.HandleFunc("PATCH /v1/worker/student/{id}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerPatchStudent)) // TODO!
	mux.HandleFunc("GET /v1/worker/application", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerGetApplications))
	mux.HandleFunc("GET /v1/worker/student/application/{id}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerGetApplicationsForStudent))
	mux.HandleFunc("PATCH /v1/worker/application/close/{id}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerCloseApplication))
	mux.HandleFunc("GET /v1/worker/application/download/{link}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerDownloadApplication))
	mux.HandleFunc("GET /v1/worker/document/download/{link}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerDownloadDocument))
	mux.HandleFunc("GET /v1/worker/student/document/{id}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerGetDocumentForUser))
	mux.HandleFunc("POST /v1/worker/document", s.SecureHandlerWithOutId(entities.UserWorker, s.WorkerAddDocument))
	mux.HandleFunc("PATCH /v1/worker/document/{id}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerPatchDocument))
	mux.HandleFunc("PATCH /v1/worker/document/file/{link}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerPatchFileDocument))
	mux.HandleFunc("DELETE /v1/worker/document/{id}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerDeleteDocument))

	mux.HandleFunc("GET /v1/student", s.SecureHandler(entities.UserStudent, s.StudentGetSelf))
	mux.HandleFunc("GET /v1/student/photo", s.SecureHandler(entities.UserStudent, s.StudentGetSelfPhoto))
	mux.HandleFunc("GET /v1/student/document", s.SecureHandler(entities.UserStudent, s.StudentGetSelfDocuments))
	mux.HandleFunc("GET /v1/student/document/download/{link}", s.SecureHandlerWithOutId(entities.UserStudent,
		s.StudentDownloadDocument))
	mux.HandleFunc("GET /v1/student/application", s.SecureHandler(entities.UserStudent, s.StudentGetApplications))
	mux.HandleFunc("POST /v1/student/application", s.SecureHandler(entities.UserStudent, s.StudentAddApplication))
	mux.HandleFunc("GET /v1/student/application/download/{link}", s.SecureHandler(entities.UserStudent,
		s.StudentDownloadApplication))

	server := http.Server{
		Addr:        fmt.Sprintf("0.0.0.0:%s", s.Port),
		IdleTimeout: s.IdleTimeout,
		Handler:     mux,
	}

	return server.ListenAndServe()
}
