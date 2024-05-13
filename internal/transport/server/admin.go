package server

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entities"
	"studentRecordsApp/internal/transport/server/httpEntity"
)

func (s Server) AdminGetSelf(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	result, err := s.user.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	httpResult, cErr := casts.UserEntitieToHttp(r.Context(), result)
	if cErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, cErr := json.Marshal(httpResult)
	if cErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) AdminPatchSelf(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	patchUser := entities.User{
		Id:        id,
		FirstName: r.PostFormValue("first_name"),
		LastName:  r.PostFormValue("last_name"),
		Surname:   r.PostFormValue("surname"),
		Email:     r.PostFormValue("email"),
	}

	if err := s.user.Update(r.Context(), patchUser); err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}
}

func (s Server) AdminGetAllWorkers(w http.ResponseWriter, r *http.Request) {
	result, err := s.user.GetAllWorker(r.Context())
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	httpResult := make([]httpEntity.User, 0, len(result))
	for idx := range result {
		tmp, err := casts.UserEntitieToHttp(r.Context(), result[idx])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		httpResult = append(httpResult, tmp)
	}

	body, cErr := json.Marshal(httpResult)
	if cErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) AdminAddWorker(w http.ResponseWriter, r *http.Request) {
	user := entities.User{
		FirstName: r.PostFormValue("first_name"),
		LastName:  r.PostFormValue("last_name"),
		Surname:   r.PostFormValue("surname"),
		Email:     r.PostFormValue("email"),
		Password:  r.PostFormValue("password"),
		Role:      r.PostFormValue("role"),
	}

	if err := s.user.Add(r.Context(), user); err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s Server) AdminPatchWorker(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	patchUser := entities.User{
		Id:        id,
		FirstName: r.PostFormValue("first_name"),
		LastName:  r.PostFormValue("last_name"),
		Surname:   r.PostFormValue("surname"),
		Email:     r.PostFormValue("email"),
	}

	if err := s.user.Update(r.Context(), patchUser); err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}
}

func (s Server) AdminDeleteWorker(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.user.Delete(r.Context(), id); err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}
}
