package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	S2   S2   `json:"s2,omitempty"`
	HTTP HTTP `json:"http,omitempty"`
}
type S2 struct {
	GrpcAddress string `json:"grpc_address,omitempty"`
}

type HTTP struct {
	Port string `json:"port,omitempty"`
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
