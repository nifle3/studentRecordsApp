package server

import (
	"net/http"
	"studentRecordsApp/internal/entites"

	"github.com/gin-gonic/gin"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/transport/server/jsonStruct"
)

func (s *Server) GetUserSelfAccount(c *gin.Context) {
	id := c.Request.Header.Get("id")
	if id == "" {
		c.JSON(400, gin.H{
			"error": "id is required",
		})
		return
	}

	result, err := s.user.Get(id, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	value := casts.UserEntitieToJson(result, c)
	c.JSON(http.StatusOK, value)
}

func (s *Server) GetAllWorker(c *gin.Context) {
	results, err := s.user.GetAllWorker(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	jsonResults := make([]jsonStruct.UserWithId, 0, len(results))

	for _, value := range results {
		jsonResults = append(jsonResults, casts.UserEntitieToJsonWithId(value, c))
	}

	c.JSON(http.StatusOK, jsonResults)
	return
}

func (s *Server) GetWorkerById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	result, err := s.user.Get(id, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	value := casts.UserEntitieToJsonWithId(result, c)
	c.JSON(http.StatusOK, value)
}

func (s *Server) UpdateWorker(c *gin.Context) {
	var user jsonStruct.UserWithId
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userResult := casts.UserJsonWithIdToEntitie(user, entities.entities.UserWorker, c)
	userResult.Role = entities.UserWorker

	err := s.user.Update(userResult, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) DeleteWorker(c *gin.Context) {
	id := c.Query("id")
	err := s.user.Delete(id, c)
	if err != nil {
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) AddWorker(c *gin.Context) {
	var user jsonStruct.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userResult := casts.UserJsonToEntitie(user, entities.UserWorker, "", c)

	err := s.user.Add(userResult, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
