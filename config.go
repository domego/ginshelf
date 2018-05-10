package main

import "github.com/koding/multiconfig"

var (
	configFile = "config/config.yaml"
	cfg        *Config
)

type (
	// Config struct
	Config struct {
		Address string `default:"0.0.0.0:8888"`
		Env     string `default:"debug"`
	}
)

func loadConfig() {
	cfg = new(Config)
	m := multiconfig.NewWithPath(configFile) // supports TOML, JSON and YAML
	m.MustLoad(cfg)
}
