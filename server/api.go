package server

import (
	"HarborMaster/managers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) buildAPIRoutes() {
	api := s.Router.Group("/api")
	api.DELETE("/:repo/:tag", s.onDeleteTagHandler())
}

func (s *Server) onDeleteTagHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := c.Param("repo")
		tag := c.Param("tag")

		err := managers.GetRegistryManager().DeleteTag(repo, tag)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Repository and tag not found or deletion disabled!"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": true})
	}
}
