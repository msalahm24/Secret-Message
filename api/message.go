package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahmoud24598salah/Secret-Message/services"
	"net/http"
)

type sendMessageReq struct {
	ToPhoneNumber string `json:"to_phone_number" binding:"required"`
	Content       string `json:"message" binding:"required"`
	Type          string `json:"type" binding:"required"`
}



func (server *Server) sendMessage(ctx *gin.Context) {
	var req sendMessageReq
	err := ctx.ShouldBindJSON(&req)
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


