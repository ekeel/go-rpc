package main

import (
	"log"
	"client/rpcclient"
)

func main() {
	testPluginClient := rpcclient.NewRPCClient("localhost", "42323", "test_plugin")
	
	log.Print(testPluginClient.ToString())
	
	response, err := testPluginClient.Call("{\"working\": \"true\"}")
	if err != nil {
		log.Fatal(err)
	}
	
	log.Print(response)
}