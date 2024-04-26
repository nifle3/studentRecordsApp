package server

import (
	"html/template"
	"net/http"

	"github.com/golang-jwt/jwt"

	"studentRecordsApp/internal/service/entites"
)

type authResponse struct {
	Email    string
	Password string
	Role     string
}

func (s *Server) auth(w http.ResponseWriter, r *http.Request) {
	result := authResponse{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Role:     r.FormValue("role"),
	}

	if result.Role == roleStudent {
		student, b, err := s.student.Login(result.Email, result.Password, r.Context())
		if err != nil || !b {
			tmp, _ := template.ParseFS(s.fs, "log.gohtml")
			tmp.Execute(w, true)
			return
		}

		setJwt(w, student.Id, result.Role)
		w.Write([]byte("Hello student"))
		return
	}

	var user entities.User
	var err error

	if result.Role == roleAdmin {
		user, err = s.user.Login(result.Password, result.Email, entities.UserAdmin, r.Context())
		if err != nil {
			tmp, _ := template.ParseFS(s.fs, "log.gohtml")
			tmp.Execute(w, true)
			w.WriteHeader(500)
			return
		}

		w.Write([]byte("hello admin"))
	} else if result.Role == roleWorker {
		user, err = s.user.Login(result.Password, result.Email, entities.UserWorker, r.Context())
		if err != nil {
			tmp, _ := template.ParseFS(s.fs, "log.gohtml")
			tmp.Execute(w, true)
			return
		}

		w.Write([]byte("hello worker"))
	} else {
		w.WriteHeader(400)
		return
	}

	setJwt(w, user.Id, result.Role)
	w.WriteHeader(200)
}

func setJwt(w http.ResponseWriter, id, role string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		Id:   id,
		Role: role,
	})

	signedString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	cookie := &http.Cookie{
		Name:     tokenCookie,
		Value:    signedString,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
}
