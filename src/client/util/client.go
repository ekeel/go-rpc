package util

import (
	"client/rpcclient"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/phayes/freeport"
)

// CreateClients returns a slice of references to RPCClient's.
func CreateClients(conf *Configuration) (map[string]*rpcclient.RPCClient, error) {
	clients := make(map[string]*rpcclient.RPCClient)

	pluginDir := fmt.Sprintf("%v", conf.Dictionary["plugin_directory"])
	root := pluginDir

	filesInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return clients, err
	}

	for _, fi := range filesInfo {
		if len(fi.Name()) > 0 {
			host := "localhost"
			name := (strings.Split(fi.Name(), "."))[0]
			port, err := freeport.GetFreePort()
			if err != nil {
				return clients, err
			}

			client := rpcclient.NewRPCClient(
				host,
				(fmt.Sprintf("%v", port)),
				path.Join(pluginDir, fi.Name()),
			)

			clients[name] = client
		}
	}

	return clients, nil
}
