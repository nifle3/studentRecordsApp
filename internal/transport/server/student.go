package server

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entities"
	"studentRecordsApp/internal/transport/server/httpEntity"
)

func (s Server) StudentGetSelf(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	result, err := s.student.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	studentSelf := casts.StudentEntitieToHttpSelf(r.Context(), result)

	jsonResult, cErr := json.Marshal(studentSelf)
	if cErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResult); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) StudentGetSelfPhoto(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	image, err := s.student.GetImageById(r.Context(), id)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	if _, err := w.Write(image); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) StudentGetSelfDocuments(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	result, err := s.document.GetAllForUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	httpResult := make([]httpEntity.DocumentSelf, 0, len(result))
	for _, value := range result {
		tmp, err := casts.DocumentEntitieToDocumentSelf(r.Context(), value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		httpResult = append(httpResult, tmp)
	}

	jsonResult, cErr := json.Marshal(httpResult)
	if cErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResult); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) StudentDownloadDocument(w http.ResponseWriter, r *http.Request) {
	link := r.PathValue("link")
	if link == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	result, err := s.document.DownloadDocument(r.Context(), link)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	if _, err := w.Write(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) StudentGetApplications(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	result, err := s.application.GetAllForUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	httpResult := make([]httpEntity.ApplicationGet, 0, len(result))
	for _, value := range result {
		tmp, err := casts.ApplicationEntitieToApplicationGet(r.Context(), value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		httpResult = append(httpResult, tmp)
	}

	jsonResult, cErr := json.Marshal(httpResult)
	if cErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResult); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) StudentAddApplication(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	file, fileInfo, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	patchResult := entities.Application{
		StudentId:   userId,
		ContactInfo: r.PostFormValue("contact_info"),
		Text:        r.PostFormValue("text"),
		Name:        r.PostFormValue("name"),
		File:        file,
	}

	if err := s.application.Add(r.Context(), patchResult, fileInfo.Size); err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s Server) StudentDownloadApplication(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	link := r.PathValue("link")
	if link == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	result, err := s.application.DownloadWithCheckId(r.Context(), link, id)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	if _, err := w.Write(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
