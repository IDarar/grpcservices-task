package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	GRPC         GRPCConfig   `json:"grpc,omitempty"`
	GoogleSearch GoogleSearch `json:"google_search,omitempty"`
}
type GRPCConfig struct {
	Port string `json:"port,omitempty"`
}

type GoogleSearch struct {
	ApiKey string `json:"api_key,omitempty"`
	Cx     string `json:"cx,omitempty"`
}

func Init(cfgPath string) (Config, error) {
	b, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}

	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
