package app

import "net/http"

func Start() {
	// storage
	// word
	// service

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, World!"))
	})
	http.ListenAndServe(":8080", nil)
}
