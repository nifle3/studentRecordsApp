package serveMux

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"studentRecordsApp/internal/service"
)

type SecureHandler func(w http.ResponseWriter, r *http.Request, userId uuid.UUID)

type AuthMux struct {
	authService service.Auth
}

const isLoginValue = "1"
const cookieLoginName = "login"
const cookieTokenName = "token"

func NewAuthMux(authService service.Auth) *AuthMux {
	return &AuthMux{
		authService: authService,
	}
}

func (a AuthMux) Login(w http.ResponseWriter, r *http.Request) {
	password := r.PostFormValue("password")
	login := r.PostFormValue("login")
	role := r.PostFormValue("role")

	auth, err := a.authService.Auth(password, login, role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	loginCookie := http.Cookie{
		Name:     cookieLoginName,
		Value:    isLoginValue,
		HttpOnly: false,
		Expires:  time.Now().Add(time.Hour * 24 * 100),
		SameSite: http.SameSiteStrictMode,
	}

	cookie := http.Cookie{
		Name:     cookieTokenName,
		Value:    auth,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 100),
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)
	http.SetCookie(w, &loginCookie)
	w.WriteHeader(http.StatusOK)
}

func (a AuthMux) SecureHandler(requireRole string, next SecureHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieLoginName)
		if err != nil {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}

		if cookie.Value != isLoginValue {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		cookie, err = r.Cookie(cookieTokenName)
		if err != nil {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}

		userId, err := a.authService.ValidateRequireRole(cookie.Value, requireRole)
		if err != nil {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r, userId)
	}
}
