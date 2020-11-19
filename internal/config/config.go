package config

import (
	"encoding/json"
	"github.com/aaroalan/mujina/internal/help"
	"io/ioutil"
	"os"
)

// Config : Struct that will parse the config JSON file.
type Config struct {
	Endpoints []Endpoint `json:"endpoints"`
}

// NewConfig : Parses a JSON file to a Config struct.
func NewConfig(path string) (Config, error) {
	f, err := os.Open(path)
	if help.HasError(&err) {
		return Config{}, err
	}
	defer f.Close()
	rawJSON, _ := ioutil.ReadAll(f)
	var config Config
	err = json.Unmarshal(rawJSON, &config)
	if help.HasError(&err) {
		return Config{}, err
	}
	return config, nil
}
