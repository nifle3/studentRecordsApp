package serveMux

import (
    "net/http"

    "github.com/google/uuid"
)

type AdminMux struct {
}

func NewAdminMux() *AdminMux {
    return &AdminMux{}
}

func (a AdminMux) GetSelf(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

}

func (a AdminMux) UpdateSelf(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

}

func (a AdminMux) GetAllWorker(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

}

func (a AdminMux) GetWorkerByID(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

}

func (a AdminMux) AddWorker(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

}

func (a AdminMux) UpdateWorker(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

}
