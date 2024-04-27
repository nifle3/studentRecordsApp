package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/transport/server/jsonStruct"
)

func (s *Server) GetSelfStudent(c *gin.Context) {
	id := c.GetHeader("id")

	get, err := s.student.Get(id, c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result := casts.StudentEntitieToJson(get, c)
	c.JSON(http.StatusOK, result)
}

func (s *Server) UpdateSelfStudent(c *gin.Context) {
	id := c.GetHeader("id")
	var result jsonStruct.StudentLong
	result.Id = id
	err := c.BindJSON(&result)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	student := casts.StudentJsonToEntitie(result, c)
	err = s.student.Update(student, c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) GetAllApplications(c *gin.Context) {
	id := c.GetHeader("id")
	applications, err := s.application.GetAllForUser(id, c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	results := make([]jsonStruct.Application, 0, len(applications))
	for _, application := range applications {
		results = append(results, casts.ApplicationEntitesToJson(application, c))
	}

	c.JSON(http.StatusOK, results)
}

func (s *Server) GetApplication(c *gin.Context) {
	id := c.Param("id")
	application, err := s.application.GetById(id, c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result := casts.ApplicationEntitesToJson(application, c)
	c.JSON(http.StatusOK, result)
}

func (s *Server) CreateApplication(c *gin.Context) {
	var result jsonStruct.ApplicationAdded
	err := c.BindJSON(&result)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	application := casts.ApplicationJsonAddedToEntites(result, c)
	err = s.application.Add(application, c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) UpdateApplication(c *gin.Context) {
	id := c.Param("id")
	var result jsonStruct.Application
	result.Id = id
	err := c.BindJSON(&result)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	application := casts.ApplicationJsonToEntites(result, c)
	err = s.application.Update(application, c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) AllDocumentsForStudent(c *gin.Context) {
	id := c.Param("id")
	documents, err := s.document.GetAllForUser(id, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	results := make([]jsonStruct.DocumentWithoutFile, 0, len(documents))
	for _, value := range documents {
		results = append(results, casts.DocumentEntiteToJsonShort(value, c))
	}

	c.JSON(http.StatusOK, results)
}

func (s *Server) GetDocument(context *gin.Context) {
	id := context.Param("id")
	userId := context.GetHeader("id")
	document, err := s.document.GetById(id, userId, context)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result := casts.DocumentEntiteToJson(document, context)
	context.JSON(http.StatusOK, result)
}
