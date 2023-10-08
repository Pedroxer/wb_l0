package api

import (
	"github.com/Pedroxer/wbL0/util"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
)

type Server struct {
	router *gin.Engine
	Con    stan.Conn
	Cache  *util.Cache
}

func NewServer(Con stan.Conn, cache *util.Cache) *Server {
	server := &Server{
		Con:   Con,
		Cache: cache,
	}
	server.setupRotes()
	return server
}
func (serv *Server) setupRotes() {
	router := gin.Default()
	router.GET("getJson", serv.getJson)
	router.POST("postJson", serv.postJson)
	serv.router = router
}

func (serv *Server) Start() error {
	return serv.router.Run(":8080")
}
