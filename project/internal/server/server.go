package server

import (
	"project/AI"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Host string
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

	engine.GET("/checkUser", r.checkHandler)

	return engine
}

func (r *Server) Start() {
	r.newApi().Run(r.Host)
}

type User struct {
	NameRecruit        string `json:"nameRecruit"`
	SpecialityRecruit  string `json:"specialityRecruit"`
	BirthDateRecruit   int    `json:"birthDateRecruit"`
	ExpirienceRecruit  int    `json:"expirienceRecruit"`
	NameEmployee       string `json:"nameEmployee"`
	SpecialityEmployee string `json:"specialityEmployee"`
	BirthDateEmployee  int    `json:"birthDateEmployee"`
	ExpirienceEmployee int    `json:"expirienceEmployee"`
}

func (r *Server) checkHandler(ctx *gin.Context) {
	var user User
	if err := json.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	ans, _ := AI.AiResponse(user.NameRecruit, user.SpecialityRecruit,
		user.BirthDateRecruit, user.ExpirienceRecruit,
		user.NameEmployee, user.SpecialityEmployee,
		user.BirthDateEmployee, user.ExpirienceEmployee)
	ctx.String(200, ans)
}
