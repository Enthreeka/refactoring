package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server struct {
		TypeServer string `yaml:"typeserver"`
		Port       string `yaml:"port"`
	} `yaml:"server"`

	Db struct {
		URL string `yaml:"url"`
	} `yaml:"db"`
}

func New(path string) (*Config, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
