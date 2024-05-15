package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"studentRecordsApp/internal/service/entities"
)

type SecureHandler func(w http.ResponseWriter, r *http.Request, userId uuid.UUID)

const cookieTokenName = "token"

func (s Server) GetRole(w http.ResponseWriter, r *http.Request) {
	role := []string{entities.UserAdmin, entities.UserWorker, entities.UserStudent}
	result, err := json.Marshal(role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) GetWorkerRole(w http.ResponseWriter, r *http.Request) {
	role := []string{entities.UserAdmin, entities.UserWorker}
	result, err := json.Marshal(role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) Login(w http.ResponseWriter, r *http.Request) {
	password := r.PostFormValue("password")
	login := r.PostFormValue("login")
	role := r.PostFormValue("role")
	if password == "" || login == "" || role == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	auth, err := s.auth.Auth(r.Context(), role, login, password)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	cookie := http.Cookie{
		Name:     cookieTokenName,
		Value:    auth,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 100),
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
}

func (s Server) CheckRole(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie(cookieTokenName)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (s Server) SecureHandler(requireRole string, next SecureHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieTokenName)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userId, cErr := s.auth.ValidateRequireRole(r.Context(), cookie.Value, requireRole)
		if cErr != nil {
			http.Error(w, cErr.GetMsg(), cErr.GetCode())
			return
		}

		next(w, r, userId)
	}
}

func (s Server) SecureHandlerWithOutId(requireRole string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieTokenName)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, cErr := s.auth.ValidateRequireRole(r.Context(), cookie.Value, requireRole)
		if cErr != nil {
			http.Error(w, cErr.GetMsg(), cErr.GetCode())
			return
		}

		next(w, r)
	}
}
