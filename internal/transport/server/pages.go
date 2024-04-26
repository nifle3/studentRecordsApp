package server

import (
	"github.com/golang-jwt/jwt"
	"html/template"
	"net/http"
)

func (s *Server) authPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err == nil {
		var redirectPage string
		withClaims, err := jwt.ParseWithClaims(cookie.Value, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil {
			w.WriteHeader(400)
			return
		}

		claims, isValid := withClaims.Claims.(*jwtClaims)
		if !isValid {
			w.WriteHeader(400)
			return
		}

		if claims.Role == roleStudent {
			redirectPage = "/student"
		}

		if claims.Role == roleAdmin {
			redirectPage = "/admin"
		}

		if claims.Role == roleWorker {
			redirectPage = "/employee"
		}

		http.Redirect(w, r, redirectPage, http.StatusSeeOther)
		return
	}

	page, err := template.ParseFS(s.fs, "log.gohtml")
	if err != nil {
		return
	}

	page.Execute(w, nil)
}
