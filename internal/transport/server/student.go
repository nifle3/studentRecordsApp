package server

import (
	"net/http"
	"studentRecordsApp/internal/casts"

	"github.com/google/uuid"
)

func (s Server) StudentGetSelf(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	_, err := s.student.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s Server) StudentGetSelfDocuments(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	_, err := s.document.GetAllForUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s Server) StudentGetDocument(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	dId, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, cErr := s.document.GetById(r.Context(), dId, id)
	if err != nil {
		http.Error(w, cErr.GetMsg(), cErr.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s Server) StudentGetApplications(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	_, err := s.application.GetAllForUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s Server) StudentGetApplication(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	appId, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, cErr := s.application.GetByIdAndUserId(r.Context(), appId, id)
	if err != nil {
		http.Error(w, cErr.GetMsg(), cErr.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}
