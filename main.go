package main

import (
	"log"

	"github.com/mahmoud24598salah/Secret-Message/api"
	"github.com/sfreiberg/gotwilio"
)

const (
	SERVER_ADDRESS   = "0.0.0.0:8080"
	TwilioAccountSid = "ACc19ea94a2e3de488564c1a3e20b76ad3"
	TwilioAuthToken  = ""
)

func main() {
	twilio := gotwilio.NewTwilioClient(TwilioAccountSid, TwilioAuthToken)
	server := api.NewServer(twilio)
	err := server.Start(SERVER_ADDRESS)

	if err != nil {
		log.Fatal("Can not start server", err)
	}
}
