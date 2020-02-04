package rpcclient

import (
	"bytes"
	"fmt"
	"net/rpc"
	"os/exec"
	"errors"
)

type RPCClient struct {
	Client *rpc.Client
	Host string
	Port string
	PluginFile string
}

// NewRPCClient returns a reference to a `RPCClient`.
func NewRPCClient(host, port, pluginFile string) *RPCClient {
	client := RPCClient{
		Host: host,
		Port: port,
		PluginFile: pluginFile,
	}
	
	return &client
}

// ToString returns the string representation of the struct.
func (client *RPCClient) ToString() string {
	return fmt.Sprintf(
		"{\"host\": \"%s\", \"port\": \"%s\", \"plugin_file\": \"%s\"}",
		client.Host,
		client.Port,
		client.PluginFile,
	)
}

// StartServer starts the related RPC route server.
func (client *RPCClient) StartServer() error {
	var err error
	
	go func(cli *RPCClient) {
		cmd := exec.Command(client.PluginFile, client.Port)
		
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		
		e := cmd.Run()
		if e != nil {
			err = e
		}
		
		if stderr.Len() > 0 {
			err = errors.New(stderr.String())
		}
	}(client)
	
	if err != nil {
		return err
	}
	
	return nil
}

// Dial calls the rpc.Client.Dial function.
func (client *RPCClient) Dial() error {
	rpcc, err := rpc.Dial("tcp", fmt.Sprintf("%s:%s", client.Host, client.Port))
	if err != nil {
		return err
	}
	
	client.Client = rpcc
	
	return nil
}

// Call calls the following workflow `StartServer` : `Dial` : `Client.Call`.
func (client *RPCClient) Call(payload string) (string, error) {
	var response string
	
	err := client.StartServer()
	if err != nil {
		return response, err
	}
	fmt.Println("***************** 1 *****************")
	
	err = client.Dial()
	if err != nil {
		return response, err
	}
	fmt.Println("***************** 2 *****************")
	
	err = client.Client.Call("Listener.Execute", payload, &response)
	if err != nil {
		return response, err
	}
	fmt.Println("***************** 3 *****************")
	
	return response, nil
}