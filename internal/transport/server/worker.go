package server

import (
	"github.com/google/uuid"
	"net/http"
	"studentRecordsApp/internal/casts"
)

// TODO: add json struct
func (s Server) WorkerGetSelf(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	_, err := s.user.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TODO;
func (s Server) WorkerGetAllStudents(w http.ResponseWriter, r *http.Request) {
	_, err := s.student.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TODO:
func (s Server) WorkerGetStudent(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, cErr := s.student.Get(r.Context(), id)
	if cErr != nil {
		http.Error(w, cErr.GetMsg(), cErr.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TODO:
func (s Server) WorkerGetStudentImage(w http.ResponseWriter, r *http.Request) {

}

// TODO: отрефакторить и student service кал туду
func (s Server) WorkerAddStudent(w http.ResponseWriter, r *http.Request) {
	image, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func() {
		if err := image.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()

	w.WriteHeader(http.StatusCreated)
}

// TODO:
func (s Server) WorkerPatchStudent(w http.ResponseWriter, r *http.Request) {

}

// TODO add json struct
func (s Server) WorkerGetApplications(w http.ResponseWriter, r *http.Request) {
	_, err := s.application.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TODO
func (s Server) WorkerGetApplicationsForStudent(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, cErr := s.application.GetAllForUser(r.Context(), id)
	if cErr != nil {
		http.Error(w, cErr.GetMsg(), cErr.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s Server) WorkerGetApplication(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, cErr := s.application.GetById(r.Context(), id)
	if cErr != nil {
		http.Error(w, cErr.GetMsg(), cErr.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s Server) WorkerCloseApplication(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if cErr := s.application.ChangeStatusToFinish(r.Context(), id); cErr != nil {
		http.Error(w, cErr.GetMsg(), cErr.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Todo
func (s Server) WorkerDownloadDocument(w http.ResponseWriter, r *http.Request) {
	_, err := casts.StringToUuid("id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s Server) WorkerGetDocumentForUser(w http.ResponseWriter, r *http.Request) {

}

func (s Server) WorkerGetDocument(w http.ResponseWriter, r *http.Request) {

}
