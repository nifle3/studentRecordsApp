package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entities"
	"studentRecordsApp/internal/transport/server/httpEntity"
)

const layoutISO = "2006-01-02"

func (s Server) WorkerGetSelf(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	result, err := s.user.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	httpResult, cErr := casts.UserEntitieToHttp(r.Context(), result)
	if cErr != nil {
		http.Error(w, cErr.Error(), http.StatusInternalServerError)
		return
	}

	body, cErr := json.Marshal(httpResult)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) WorkerGetAllStudents(w http.ResponseWriter, r *http.Request) {
	result, err := s.student.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	jsonResult := make([]httpEntity.StudentShort, 0, len(result))
	for _, value := range result {
		tmp, err := casts.StudentEntitieToStudentShort(r.Context(), value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResult = append(jsonResult, tmp)
	}

	body, cErr := json.Marshal(jsonResult)
	if cErr != nil {
		http.Error(w, cErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) WorkerGetStudent(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, cErr := s.student.Get(r.Context(), id)
	if cErr != nil {
		http.Error(w, cErr.GetMsg(), cErr.GetCode())
		return
	}

	httpResult, err := casts.StudentEntitieToHttp(r.Context(), result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(httpResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) WorkerGetStudentImage(w http.ResponseWriter, r *http.Request) {
	link := r.PathValue("link")
	image, err := s.student.GetImage(r.Context(), link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	if _, err := w.Write(image); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) WorkerAddStudent(w http.ResponseWriter, r *http.Request) {
	image, imageInfo, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func() {
		if err := image.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()

	group, err := strconv.Atoi(r.FormValue("group"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	passportSeria, err := strconv.Atoi(r.FormValue("passport_seria"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	passportNumber, err := strconv.Atoi(r.FormValue("passport_number"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	birth, err := time.Parse(layoutISO, r.FormValue("birthdate"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hostNumber, err := strconv.Atoi(r.FormValue("house"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	apartment, err := strconv.Atoi(r.FormValue("apartment"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	course, err := strconv.Atoi(r.PostFormValue("course"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := entities.Student{
		FirstName:       r.FormValue("name"),
		Surname:         r.FormValue("surname"),
		LastName:        r.FormValue("last_name"),
		Group:           group,
		PassportSeria:   passportSeria,
		PassportNumber:  passportNumber,
		BirthDate:       birth,
		Email:           r.FormValue("email"),
		Password:        r.FormValue("password"),
		Photo:           image,
		Country:         r.FormValue("country"),
		City:            r.FormValue("city"),
		Street:          r.FormValue("street"),
		HouseNumber:     hostNumber,
		ApartmentNumber: apartment,
		Specialization:  r.FormValue("specialization"),
		Course:          course,
		PhoneNumbers:    make([]entities.PhoneNumber, 0, 0),
	}

	phoneName := "phone"
	descriptionName := "description"
	i := 1
	for r.PostForm.Has(phoneName+strconv.Itoa(i)) && r.PostForm.Has(descriptionName+strconv.Itoa(i)) {
		result.PhoneNumbers = append(result.PhoneNumbers, entities.PhoneNumber{
			Phone:       r.PostFormValue(phoneName + strconv.Itoa(i)),
			Description: r.PostFormValue(descriptionName + strconv.Itoa(i)),
		})

		i++
	}

	if err := s.student.Add(r.Context(), result, imageInfo.Size); err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// TODO:
func (s Server) WorkerPatchStudent(w http.ResponseWriter, r *http.Request) {

}

func (s Server) WorkerGetApplications(w http.ResponseWriter, r *http.Request) {
	result, err := s.application.GetAllWithInfo(r.Context())
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	jsonResult := make([]httpEntity.ApplicationWithInfo, 0, len(result))
	for idx := range result {
		tmp, err := casts.ApplicationWithInfoEntitieToHttp(r.Context(), result[idx])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResult = append(jsonResult, tmp)
	}

	jsonBody, cErr := json.Marshal(jsonResult)
	if cErr != nil {
		http.Error(w, cErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) WorkerGetApplicationsForStudent(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, cErr := s.application.GetAllForUser(r.Context(), id)
	if cErr != nil {
		http.Error(w, cErr.GetMsg(), cErr.GetCode())
		return
	}

	jsonResult := make([]httpEntity.ApplicationGet, 0, len(result))
	for idx := range result {
		tmp, err := casts.ApplicationEntitieToApplicationGet(r.Context(), result[idx])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResult = append(jsonResult, tmp)
	}

	body, err := json.Marshal(jsonResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) WorkerGetApplication(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, cErr := s.application.GetById(r.Context(), id)
	if cErr != nil {
		http.Error(w, cErr.GetMsg(), cErr.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s Server) WorkerCloseApplication(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if cErr := s.application.ChangeStatusToFinish(r.Context(), id); cErr != nil {
		http.Error(w, cErr.GetMsg(), cErr.GetCode())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s Server) WorkerDownloadDocument(w http.ResponseWriter, r *http.Request) {
	link := r.PathValue("link")
	body, err := s.document.DownloadDocument(r.Context(), link)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	if _, err := w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) WorkerDownloadApplication(w http.ResponseWriter, r *http.Request) {
	link := r.PathValue("link")
	if link == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	result, err := s.application.Download(r.Context(), link)
	if err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	if _, err := w.Write(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) WorkerGetDocumentForUser(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, cErr := s.document.GetAllForUser(r.Context(), id)
	if cErr != nil {
		http.Error(w, cErr.GetMsg(), cErr.GetCode())
		return
	}

	jsonResult := make([]httpEntity.DocumentSelf, 0, len(result))
	for idx := range result {
		tmp, err := casts.DocumentEntitieToDocumentSelf(r.Context(), result[idx])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResult = append(jsonResult, tmp)
	}

	body, err := json.Marshal(jsonResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) WorkerAddDocument(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PostFormValue("student_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, fileInfo, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	result := entities.Document{
		StudentId: id,
		Name:      r.PostFormValue("name"),
		Type:      r.PostFormValue("type"),
		File:      file,
	}

	if err := s.document.Add(r.Context(), result, fileInfo.Size); err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s Server) WorkerPatchDocument(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := entities.Document{
		Id:   id,
		Name: r.PostFormValue("name"),
		Type: r.PostFormValue("type"),
	}

	if err := s.document.Update(r.Context(), result); err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
	}
}

func (s Server) WorkerDeleteDocument(w http.ResponseWriter, r *http.Request) {
	id, err := casts.StringToUuid(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.document.Delete(r.Context(), id, r.PathValue("id")); err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}
}

func (s Server) WorkerPatchFileDocument(w http.ResponseWriter, r *http.Request) {
	link := r.PathValue("link")
	file, fileInfo, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := s.document.UpdateFile(r.Context(), link, fileInfo.Size, file); err != nil {
		http.Error(w, err.GetMsg(), err.GetCode())
		return
	}
}
