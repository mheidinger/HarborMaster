package server

import (
	"html/template"
	"net/http"
	"strings"

	"HarborMaster/managers"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func (s *Server) parseTemplates() {
	funcMap := template.FuncMap{"mod": s.templateMod}

	r := multitemplate.NewRenderer()
	r.AddFromFilesFuncs("dashboard", funcMap, "templates/base.html", "templates/dashboard.html")
	r.AddFromFiles("notfound", "templates/base.html", "templates/notfound.html")
	r.AddFromFiles("unauthorized", "templates/base.html", "templates/unauthorized.html")
	s.Router.HTMLRender = r
}

func (s *Server) templateMod(x, y, z int) bool {
	return x%y == z
}

func (s *Server) buildUIRoutes() {
	s.Router.GET("/", s.dashboardUIHandler())
	s.Router.NoRoute(s.notfoundUIHandler())
}

func (s *Server) dashboardUIHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		info, err := managers.GetRegistryManager().GetRegistryInfo()
		if err != nil {
			c.Redirect(http.StatusFound, s.getRedirectURL(*c.Request.URL, "/", &errorInfo, nil))
			return
		}

		c.HTML(http.StatusOK, "dashboard", gin.H{
			"error":   c.Query(errorURLParam),
			"success": c.Query(successURLParam),
			"info":    info,
		})
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
