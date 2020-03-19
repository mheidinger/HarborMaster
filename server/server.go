package server

import (
	"github.com/gin-gonic/gin"
)

const (
	errorURLParam   = "error"
	successURLParam = "success"
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
