package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sfreiberg/gotwilio"
)

type Server struct {
	router *gin.Engine
	twilio *gotwilio.Twilio
}

// creating new http server and setup eouting
func NewServer(twilio *gotwilio.Twilio) *Server {
	server := &Server{twilio: twilio}
	router := gin.Default()

	//add routes here

	router.POST("/messages", server.sendMessage)
	router.OPTIONS("/messages",server.optionMessage)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
