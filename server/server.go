package server

import (
	"net/url"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

const (
	errorURLParam   = "error"
	successURLParam = "success"
)

var errorInfo = "Failed to get info from registry"

type Server struct {
	Router       *gin.Engine
	neededHeader string
}

func NewServer(neededHeader string) *Server {
	s := &Server{
		Router:       gin.New(),
		neededHeader: neededHeader,
	}
	s.parseTemplates()
	s.buildRoutes()
	return s
}

func (s *Server) buildRoutes() {
	s.Router.Use(ginlogrus.Logger(log.New()), gin.Recovery())
	if s.neededHeader != "" {
		s.Router.Use(NeededHeaderMiddleware(s.neededHeader))
	}

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
