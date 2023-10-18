package http

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Host   string
	Port   string
	Routes *gin.Engine
}

func NewServer(host string, port string) *Server {
	return &Server{
		Host:   host,
		Port:   port,
		Routes: gin.Default(),
	}
}

func (s *Server) Start() error {
	return s.Routes.Run(s.Host + ":" + s.Port)
}
