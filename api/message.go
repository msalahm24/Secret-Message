package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahmoud24598salah/Secret-Message/services"
	"github.com/sfreiberg/gotwilio"
	"net/http"
)

type sendMessageReq struct {
	ToPhoneNumber string `json:"to_phone_number" binding:"required"`
	Content       string `json:"message" binding:"required"`
	Type          string `json:"type" binding:"required"`
}

const (
	TwilioAccountSid = "ACc19ea94a2e3de488564c1a3e20b76ad3"
	TwilioAuthToken  = "131ef8b7ad6209cefa34b5475339966d"
)

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
	twilio := gotwilio.NewTwilioClient(TwilioAccountSid, TwilioAuthToken)
	err, resp := services.SendingMessage(parm, twilio)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}


