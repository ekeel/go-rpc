package main

import (
	"client/util"
	"log"
)

func main() {
	conf := util.Configuration{}
	conf.LoadConfig()

	clientDict, err := util.CreateClients(&conf)
	if err != nil {
		log.Fatal(err)
	}

	response, err := clientDict["testplugin"].Call("{\"data\": \"payload_data\"}")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(response)
}
