package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/transport/server/jsonStruct"
)

func (s *Server) GetAllStudent(c *gin.Context) {
	results, err := s.student.GetAll(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	jsonResults := make([]jsonStruct.StudentShort, 0, len(results))
	for _, result := range results {
		jsonResults = append(jsonResults, casts.StudentEntitieToJsonShort(result, c))
	}

	c.JSON(200, jsonResults)
}

func (s *Server) GetStudentById(c *gin.Context) {
	id := c.Param("id")
	result, err := s.student.Get(id, c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	jsonResult := casts.StudentEntitieToJson(result, c)
	c.JSON(200, jsonResult)
}

func (s *Server) AddStudent(c *gin.Context) {
	var student jsonStruct.StudentLongWithoutLink
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	result := casts.StudentJsonWithoutLinkToEntitie(student, c)

	if err := s.student.Add(result, c); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) UpdateStudent(c *gin.Context) {
	var result jsonStruct.StudentLong
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	student := casts.StudentJsonToEntitie(result, c)
	if err := s.student.Update(student, c); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	err := s.student.Delete(id, c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})

}

func (s *Server) AddPhoneNumber(c *gin.Context) {
	userId := c.Param("id")
	var phoneNumber jsonStruct.PhoneNumberWithoutId
	if err := c.ShouldBindJSON(&phoneNumber); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	phoneNumber.StudentId = userId

	result := casts.PhoneNumberJsonToEntities(phoneNumber, c)
	if err := s.phoneNumber.Add(result, c); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) AddDocument(c *gin.Context) {
	userId := c.Param("id")
	var document jsonStruct.DocumentForAdded
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	document.StudentId = userId

	result := casts.DocumentForAddedToEntite(document, c)
	if err := s.document.Add(result, c); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) DeletePhoneNumber(c *gin.Context) {
	id := c.Param("id")
	studentId := c.Param("id")
	err := s.phoneNumber.Delete(id, studentId, c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func (s *Server) UpdatePhoneNumber(c *gin.Context) {
	var phoneNumber jsonStruct.PhoneNumber
	if err := c.ShouldBindJSON(&phoneNumber); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	result := casts.PhoneNumberJsonLongToEntities(phoneNumber, c)
	if err := s.phoneNumber.Update(result, c); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) DeleteDocument(c *gin.Context) {
	id := c.Param("documentId")
	userId := c.Param("id")

	if err := s.document.Delete(id, userId, c); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) GetPhoneNumbers(c *gin.Context) {
	userId := c.Param("id")
	results, err := s.phoneNumber.GetAllForUser(userId, c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	jsonResults := make([]jsonStruct.PhoneNumber, 0, len(results))
	for _, result := range results {
		jsonResults = append(jsonResults, casts.PhoneNumberEntitiesToJson(result, c))
	}

	c.JSON(200, jsonResults)

}

func (s *Server) GetDocuments(c *gin.Context) {
	userId := c.Param("id")

	documentations, err := s.document.GetAllForUser(userId, c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	results := make([]jsonStruct.DocumentWithoutFile, 0, len(documentations))
	for _, value := range documentations {
		results = append(results, casts.DocumentEntiteToJsonShort(value, c))
	}

	c.JSON(200, results)
}

func (s *Server) GetDocumentById(c *gin.Context) {
	id := c.Param("documentId")
	userId := c.Param("id")
	result, err := s.document.GetById(id, userId, c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	jsonResult := casts.DocumentEntiteToJson(result, c)

	c.JSON(200, jsonResult)
}

func (s *Server) GetApplications(c *gin.Context) {

}

func (s *Server) GetApplicationById(c *gin.Context) {

}

func (s *Server) CloseApplication(c *gin.Context) {

}
