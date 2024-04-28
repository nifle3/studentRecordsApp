package server

import (
	"log"
	"net/http"
	"studentRecordsApp/internal/entites"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type authResponse struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

func (s *Server) auth(c *gin.Context) {
	var response authResponse
	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Printf("%#v", response)

	if response.Role == roleStudent {
		student, err := s.student.Login(response.Email, response.Password, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		setJwt(c, student.Id, response.Role)
		c.Status(http.StatusOK)
		return
	}

	if response.Role == roleAdmin {
		user, err := s.user.Login(response.Password, response.Email, entities.entities.UserAdmin, c)
		if err != nil {
			log.Printf("%s", err.Error())
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		setJwt(c, user.Id, response.Role)
		c.Status(http.StatusOK)
		return
	}

	if response.Role == roleWorker {
		user, err := s.user.Login(response.Password, response.Email, entities.UserWorker, c)
		if err != nil {
			log.Printf("%s", err.Error())
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		setJwt(c, user.Id, response.Role)
		c.Status(http.StatusOK)
		return
	}
}

func setJwt(c *gin.Context, id, role string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		Id:   id,
		Role: role,
	})

	signedString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		c.Status(500)
		return
	}

	c.SetCookie(tokenCookie, signedString, 3600, "/", "localhost", false, true)
}
