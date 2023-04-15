package conf

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	PcapFilePath string `json:"pcapFilePath"`
}

func ReadConfig() (*Config, error) {
	const confFile = "config.json"

	conf := new(Config)

	value, err := ioutil.ReadFile(confFile)
	if err != nil {
		return conf, err
	}

	// decode json
	err = json.Unmarshal([]byte(value), conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
