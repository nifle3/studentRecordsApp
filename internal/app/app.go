package app

import (
	"os"
	"studentRecordsApp/internal/transport/server"
)

func Start() {
	// storage
	// word
	// service
	err := server.Start()
	if err != nil {
		os.Exit(1)
	}
}
