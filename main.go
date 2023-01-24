package main

import (
	"log"
	"github.com/mahmoud24598salah/Secret-Message/api"
	"github.com/mahmoud24598salah/Secret-Message/util"
	"github.com/sfreiberg/gotwilio"
)

func main() {
	config,err := util.LoadConfig(".")

	if err != nil{
		log.Fatal("Can not read the config file",err)
	}

	twilio := gotwilio.NewTwilioClient(config.TWILIO_ACCOUNT_SID, config.TWILIO_AUTH_TOKEN)
	server := api.NewServer(twilio)
	err = server.Start(config.SreverAddress)

	if err != nil {
		log.Fatal("Can not start server", err)
	}
}
