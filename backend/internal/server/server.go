package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Host string
}

type User struct {
	Name     string
	Surname  string
	Password int
}

func New(host string) *Server {
	s := &Server{
		Host: host,
	}
	return s
}

func (r *Server) newApi() *gin.Engine {
	engine := gin.New()
	engine.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	return engine
}

func (r *Server) Start() {
	r.newApi().Run(r.Host)
}
