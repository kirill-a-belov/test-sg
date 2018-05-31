package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/kirill-a-belov/test-sg/models"
	"github.com/kirill-a-belov/test-sg/storages/pg"
)

type Server struct {
	config  *ServerConfig
	storage *pg.PGStorage
}

func NewServer(config *ServerConfig, storage *pg.PGStorage) *Server {
	return &Server{
		config:  config,
		storage: storage,
	}
}

func (s *Server) Serve() error {
	server := gin.Default()

	// Routing
	server.POST("/by_url", s.NewByURLHandler())

	server.GET("/by_pid/:pid", s.NewByPIDHandler())

	addr := fmt.Sprintf("%v:%v", s.config.Address, s.config.Port)

	return server.Run(addr)
}

func sendError(conn *net.TCPConn, code int, msg string) {
	e := &models.Error{
		Code:    code,
		Message: msg,
	}

	response, err := json.Marshal(e)
	if err != nil {
		log.Printf("error while marshalling error message : %v /n", err)
		return
	}

	conn.Write(response)
}
