package server

import "net/http"

var jwtSecretKey = []byte("very-secret-key")

type authResponse struct {
    Email    string `form:"email"`
    Password string `form:"password"`
    Role     string `form:"role"`
}

func (s *Server) Auth(w http.ResponseWriter, r *http.Request) {

}
