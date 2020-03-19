package server

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func (s *Server) parseTemplates() {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("dashboard", "templates/base.html", "templates/dashboard.html")
	r.AddFromFiles("notfound", "templates/base.html", "templates/notfound.html")
	s.Router.HTMLRender = r
}

func (s *Server) buildUIRoutes() {
	s.Router.GET("/", s.dashboardUIHandler())
	s.Router.NoRoute(s.notfoundUIHandler())
}

func (s *Server) dashboardUIHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard", gin.H{})
	}
}

func (s *Server) notfoundUIHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.UserAgent(), "Mozilla") {
			c.HTML(http.StatusNotFound, "notfound", gin.H{
				"error":   c.Query(errorURLParam),
				"success": c.Query(successURLParam),
			})
		} else {
			c.String(http.StatusNotFound, "Not Found", gin.H{})
		}
	}
}
