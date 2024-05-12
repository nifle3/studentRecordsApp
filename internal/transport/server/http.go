package server

import (
	"fmt"
	"net/http"
	entities "studentRecordsApp/internal/service/entites"
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

	// WORK
	mux.HandleFunc("POST /v1/auth", s.Login)
	mux.HandleFunc("GET /v1/role", s.GetRole)
	mux.HandleFunc("GET /v1/role/worker", s.GetWorkerRole)

	// WORK
	mux.HandleFunc("GET /v1/admin", s.SecureHandler(entities.UserAdmin, s.AdminGetSelf))
	mux.HandleFunc("PATCH /v1/admin", s.SecureHandler(entities.UserAdmin, s.AdminPatchSelf))
	mux.HandleFunc("GET /v1/admin/worker", s.SecureHandlerWithOutId(entities.UserAdmin, s.AdminGetAllWorkers))
	mux.HandleFunc("POST /v1/admin/worker", s.SecureHandlerWithOutId(entities.UserAdmin, s.AdminAddWorker))
	mux.HandleFunc("PATCH /v1/admin/worker/{id}", s.SecureHandlerWithOutId(entities.UserAdmin, s.AdminPatchWorker))
	mux.HandleFunc("DELETE /v1/admin/worker/{id}", s.SecureHandlerWithOutId(entities.UserAdmin, s.AdminDeleteWorker))

	mux.HandleFunc("GET /v1/worker", s.SecureHandler(entities.UserWorker, s.WorkerGetSelf))
	mux.HandleFunc("GET /v1/worker/student", s.SecureHandlerWithOutId(entities.UserWorker, s.WorkerGetAllStudents))
	mux.HandleFunc("GET /v1/worker/student/{id}", s.SecureHandlerWithOutId(entities.UserWorker, s.WorkerGetStudent))
	//mux.HandleFunc("GET /v1/worker/student/{id}/image", nil) // TODO!
	mux.HandleFunc("POST /v1/worker/student", s.SecureHandlerWithOutId(entities.UserWorker, s.WorkerAddStudent))
	mux.HandleFunc("PATCH /v1/worker/student/{id}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerPatchStudent))
	mux.HandleFunc("GET /v1/worker/application", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerGetApplications))
	mux.HandleFunc("GET /v1/worker/student/application/{id}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerGetApplicationsForStudent))
	mux.HandleFunc("GET /v1/worker/application/{id}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerGetApplication))
	mux.HandleFunc("PATCH /v1/worker/application/{id}/close", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerCloseApplication))
	mux.HandleFunc("GET /v1/worker/document/{id}/download", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerDownloadDocument))
	mux.HandleFunc("GET /v1/worker/student/document/{id}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerGetDocumentForUser))
	mux.HandleFunc("GET /v1/worker/document/{id}", s.SecureHandlerWithOutId(entities.UserWorker,
		s.WorkerGetDocument))
	//mux.HandleFunc("POST /v1/worker/document", nil)       //TODO!
	//mux.HandleFunc("PATCH /v1/worker/document/{id}", nil) //TODO!

	mux.HandleFunc("GET /v1/student", s.SecureHandler(entities.UserStudent, s.StudentGetSelf))
	mux.HandleFunc("GET /v1/student/document", s.SecureHandler(entities.UserStudent, s.StudentGetSelfDocuments))
	mux.HandleFunc("GET /v1/student/document/{id}", s.SecureHandler(entities.UserStudent, s.StudentGetDocument))
	//mux.HandleFunc("GET /v1/student/document/download/{id}", s.SecureHandler(entities.UserStudent, nil)) //TODO!
	mux.HandleFunc("GET /v1/student/application", s.SecureHandler(entities.UserStudent, s.StudentGetApplications))
	mux.HandleFunc("GET /v1/student/application/{id}", s.SecureHandler(entities.UserStudent, s.StudentGetApplications))
	//mux.HandleFunc("PATCH /v1/student/application/{id}", s.SecureHandler(entities.UserStudent, nil)) // TODO
	//mux.HandleFunc("POST /v1/student/application", s.SecureHandler(entities.UserStudent, nil))       // TODO

	server := http.Server{
		Addr:        fmt.Sprintf("0.0.0.0:%s", s.Port),
		IdleTimeout: s.IdleTimeout,
		Handler:     mux,
	}

	return server.ListenAndServe()
}
