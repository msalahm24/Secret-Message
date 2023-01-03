package main

import (
	"log"

	"github.com/mahmoud24598salah/Secret-Message/api"
)

const SERVER_ADDRESS = "0.0.0.0:8080"

func main(){
	server := api.NewServer()
	err := server.Start(SERVER_ADDRESS)

	if err != nil{
		log.Fatal("Can not start server",err )
	}
}