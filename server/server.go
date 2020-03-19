package server

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	errorURLParam   = "error"
	successURLParam = "success"
)

var (
	errorInfo = "Failed to get info from registry"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	s := &Server{
		Router: gin.New(),
	}
	s.parseTemplates()
	s.buildRoutes()
	return s
}

func (s *Server) buildRoutes() {
	s.Router.Use(gin.Recovery())

	s.buildUIRoutes()
	s.buildAPIRoutes()
	s.Router.Static("/static", "./static")
}

func (s *Server) getRedirectURL(reqURL url.URL, path string, errVal, successVal *string) string {
	redirectURL := reqURL
	redirectURL.Path = path
	q := redirectURL.Query()
	if errVal != nil {
		q.Set(errorURLParam, *errVal)
	} else {
		q.Del(errorURLParam)
	}
	if successVal != nil {
		q.Set(successURLParam, *successVal)
	} else {
		q.Del(successURLParam)
	}
	redirectURL.RawQuery = q.Encode()

	return redirectURL.String()
}
