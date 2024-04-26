package app

import (
    "context"
    "log"
    "os"

    "studentRecordsApp/internal/config"
    "studentRecordsApp/internal/service"
    "studentRecordsApp/internal/storage/minio"
    "studentRecordsApp/internal/storage/sql"
    "studentRecordsApp/internal/transport/server"
)

func Start() {
    cfg := config.GetConfig()

    storage, err := sql.New(cfg.GetDbConnectionString())
    if err != nil {
        log.Printf("%s", err.Error())
        os.Exit(1)
    }

    fsStorage, err := minio.New(cfg.FsEndPoint, cfg.FsPassword, cfg.FsUser, false, context.Background())
    if err != nil {
        log.Printf("%s", err.Error())
        os.Exit(1)
    }

    documentService := service.NewDocument(storage, fsStorage)
    userService := service.NewUser(storage)
    studentService := service.NewStudent(storage, fsStorage)
    applicationService := service.NewApplication(storage, fsStorage)
    phoneService := service.NewPhoneNumber(storage)

    httpServer := server.New(applicationService, studentService, phoneService, documentService, userService, nil)
    log.Fatal(httpServer.Start())

}
