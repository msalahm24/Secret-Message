package api

import (
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mahmoud24598salah/Secret-Message/services"
)

type sendMessageReq struct {
	ToPhoneNumber string `json:"to_phone_number" binding:"required"`
	Content       string `json:"message" binding:"required"`
	Type          string `json:"type" binding:"required"`
}

func (server *Server) sendMessage(ctx *gin.Context) {
	var req sendMessageReq
	err := ctx.ShouldBindJSON(&req)
	server.router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
        AllowMethods:     []string{"POST", "OPTIONS"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	parm := services.Message{
		Content: req.Content,
		To:      req.ToPhoneNumber,
		Type:    req.Type,
	}
	err, resp := services.SendingMessage(parm, server.twilio)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) optionMessage(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	ctx.Header("Access-Control-Allow-Headers", "authorization")
}
