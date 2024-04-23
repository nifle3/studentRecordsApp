package server

import (
	"github.com/google/uuid"
	"net/http"
)

func Start() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		_, _ = writer.Write([]byte(uuid.New().String()))
	})

	return http.ListenAndServe(":8080", nil)
}
