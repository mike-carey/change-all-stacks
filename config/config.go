/**
 *
 */
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

type Configs map[string]*cfclient.Config

func LoadConfigFromFile(configPath string) (Configs, error) {
	byteValue, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Configs{}, err
	}

	return LoadConfigFromBytes(byteValue)
}
// Loads the configuration and returns a configuration object
func LoadConfigFromString(jsonStr string) (Configs, error) {
	return LoadConfigFromBytes([]byte(jsonStr))
}

func LoadConfigFromBytes(byteValue []byte) (Configs, error) {
	var config Configs
	err := json.Unmarshal(byteValue, &config)
	if err != nil {
		return Configs{}, err
	}

	return config, nil
}

func (f Configs) Get(name string) (*cfclient.Config, error) {
	if c, ok := f[name]; ok {
		return c, nil
	}

	return &cfclient.Config{}, errors.New(fmt.Sprintf("Foundation not found in config: %s", name))
}
