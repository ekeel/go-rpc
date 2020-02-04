package util

import (
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// RPCConfigFileVar is the key of the environment variable containing the path to the configuration file.
const RPCConfigFileVar = "RPC_CONFIG_FILE"

// Configuration holds the parsed contents from the configuration file.
// var Configuration map[string]interface{}
//
type Configuration struct {
	Dictionary map[string]interface{}
}

func envConfig() map[string]interface{} {
	var dict map[string]interface{}

	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		dict[pair[0]] = pair[1]
	}

	return dict
}

func fileConfig(filepath string) (map[string]interface{}, error) {
	var dict map[string]interface{}

	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return dict, err
	}

	err = yaml.Unmarshal(content, &dict)
	if err != nil {
		return dict, err
	}

	return dict, nil
}

func (conf *Configuration) update(dict map[string]interface{}) {
	for k, v := range dict {
		conf.Dictionary[k] = v
	}
}

// LoadConfig reads the conf.yaml file. If the RPC_CONFIG_FILE environment variable is set, load the configuration sepecified.
func (conf *Configuration) LoadConfig() error {
	var err error
	var fileDict map[string]interface{}

	if len(os.Getenv(RPCConfigFileVar)) > 0 {
		pair := strings.SplitN(os.Getenv(RPCConfigFileVar), "=", 2)

		fileDict, err = fileConfig(pair[1])
		if err != nil {
			return err
		}

		envDict := envConfig()

		conf.update(fileDict)
		conf.update(envDict)
	}

	return nil
}
