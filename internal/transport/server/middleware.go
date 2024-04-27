package server

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/gommon/log"
	"net/http"
)

func (s *Server) authMiddleware(requireRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		token, err := jwt.ParseWithClaims(cookie, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil {
			log.Errorf("%s", err.Error())
			c.JSON(http.StatusBadRequest, err)
			return
		}

		claims, ok := token.Claims.(*jwtClaims)
		if !ok {
			log.Errorf("%s", err.Error())
			c.Status(http.StatusBadRequest)
			return
		}

		if claims.Role != requireRole {
			c.Status(http.StatusUnauthorized)
			return
		}

		c.Request.Header.Add("id", claims.Id)
		c.Request.Header.Add("role", claims.Role)

		c.Next()
	}
}
