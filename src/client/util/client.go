package util

import (
	"client/rpcclient"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/phayes/freeport"
)

// CreateClients returns a slice of references to RPCClient's.
func CreateClients(conf *Configuration) (map[string]*rpcclient.RPCClient, error) {
	var clients map[string]*rpcclient.RPCClient
	var files []string

	root := fmt.Sprintf("%v", conf.Dictionary["plugin_directory"])

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return clients, err
	}

	for _, v := range files {
		host := "localhost"
		name := strings.ReplaceAll(v, ".go", "")
		port, err := freeport.GetFreePort()
		if err != nil {
			return clients, err
		}

		clients[name] = rpcclient.NewRPCClient(host, fmt.Sprintf("%v", port), name)
	}

	return clients, nil
}
