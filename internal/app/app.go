package app

import (
	"context"
	"embed"
	"log"
	"os"

	"studentRecordsApp/internal/config"
	"studentRecordsApp/internal/service"
	"studentRecordsApp/internal/storage/minio"
	"studentRecordsApp/internal/storage/sql"
	"studentRecordsApp/internal/transport/server"
)

//go:embed web
var frontEnd embed.FS

func Start() {
	cfg := config.GetConfig()
	log.Printf("cfg is initialize")
	log.Printf("%s", cfg.GetDbConnectionString())
	storage, err := sql.New(cfg.GetDbConnectionString())
	if err != nil {
		log.Printf("%s", err.Error())
		os.Exit(1)
	}
	log.Printf("db is initialize")

	fsStorage, err := minio.New(cfg.FsEndPoint, cfg.FsPassword, cfg.FsUser, false, context.Background())
	if err != nil {
		log.Printf("%s", err.Error())
		os.Exit(1)
	}
	log.Printf("fs is initialize")

	documentService := service.NewDocument(storage, fsStorage)
	userService := service.NewUser(storage)
	studentService := service.NewStudent(storage, fsStorage)
	applicationService := service.NewApplication(storage, fsStorage)
	phoneService := service.NewPhoneNumber(storage)
	log.Printf("services is initialize")

	httpServer := server.New(applicationService, studentService, phoneService, documentService, userService, &frontEnd)
	log.Printf("Server is listening on port %s", cfg.DbPort)
	log.Fatal(httpServer.Start())
}
