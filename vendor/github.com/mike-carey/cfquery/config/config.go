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

type Foundations map[string]*cfclient.Config

// Loads the configuration and returns a configuration object
func LoadConfig(configPath string) (Foundations, error) {
	byteValue, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Foundations{}, err
	}

	var config Foundations
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return Foundations{}, err
	}

	return config, nil
}

func (f Foundations) Get(name string) (*cfclient.Config, error) {
	if c, ok := f[name]; ok {
		return c, nil
	}

	return &cfclient.Config{}, errors.New(fmt.Sprintf("Foundation not found in config: %s", name))
}
