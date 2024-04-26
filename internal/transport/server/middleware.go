package server

import (
	"github.com/golang-jwt/jwt"
	"net/http"
)

func (s *Server) authMiddleware(next http.HandlerFunc, requireRole string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err == nil {
			w.WriteHeader(400)
		}

		withClaims, err := jwt.ParseWithClaims(cookie.Value, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		value, ok := withClaims.Claims.(jwtClaims)
		if !ok {
			w.WriteHeader(400)
			return
		}

		if value.Role != requireRole {
			w.WriteHeader(403)
			return
		}

		r.Header.Add("user-id", value.Id)

		next(w, r)
	}
}
