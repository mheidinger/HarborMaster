package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) buildAPIRoutes() {
	api := s.Router.Group("/api")
	api.DELETE("/:repo/:tag", s.onDeleteTagHandler())
}

func (s *Server) onDeleteTagHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error"})
	}
}
