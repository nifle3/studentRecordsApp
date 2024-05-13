package app

import (
	"context"
	"log"
	"time"

	"studentRecordsApp/internal/config"
	"studentRecordsApp/internal/service"
	"studentRecordsApp/internal/storage/minio"
	"studentRecordsApp/internal/storage/sql"
	"studentRecordsApp/internal/transport/server"
	"studentRecordsApp/pkg/storage/db"
	"studentRecordsApp/pkg/storage/objectStorage"
)

func Start() {
	cfg := config.GetConfig()
	log.Printf("cfg is initialize")
	log.Printf("%s", cfg.GetDbConnectionString())
	ctx := context.Background()
	sqlConn := db.MustNewSqlConnection(ctx, cfg.GetDbConnectionString())
	log.Printf("db is initialize")
	minioConn := objectStorage.MustGetInstance(ctx, cfg.FsEndPoint, cfg.FsPassword, cfg.FsUser)

	documentStorage := sql.NewDocument(sqlConn)
	applicationStorage := sql.NewApplication(sqlConn)
	studentStorage := sql.NewStudent(sqlConn)
	userStorage := sql.NewUser(sqlConn)
	phoneStorage := sql.NewPhone(sqlConn)
	log.Printf("storages is initialize")

	documentFs := minio.MustNewDocument(ctx, minioConn)
	studebtPhotoFs := minio.MustNewStudentPhoto(ctx, minioConn)
	applicationFs := minio.MustNewApplication(ctx, minioConn)
	log.Printf("fs is initialize")

	documentService := service.NewDocument(documentStorage, documentFs)
	userService := service.NewUser(userStorage)
	loginService := service.NewAuth(nil, studentStorage, userStorage)
	studentService := service.NewStudent(studentStorage, studebtPhotoFs, phoneStorage)
	applicationService := service.NewApplication(applicationStorage, applicationFs)
	log.Printf("services is initialize")

	opts := server.NewOtps(cfg.ServerPort, time.Second*30)
	mux := server.NewMux(applicationService, documentService, loginService, studentService, userService)

	httpServer := server.New(opts, mux)
	log.Printf("Server is listening on port %s", cfg.ServerPort)
	log.Fatal(httpServer.Start())
}
