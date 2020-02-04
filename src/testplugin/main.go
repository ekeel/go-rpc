package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
)

const (
	// NAME contains the name of the plugin.
	NAME = "testPlugin"
)

// Listener contains the execute function for the plugin.
type Listener int

// Execute execute the primary plugin function.
func (listener *Listener) Execute(payload string, response *string) error {
	*response = payload

	return nil
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%s", os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}
