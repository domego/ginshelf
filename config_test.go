package main

import "testing"

func TestLoadConfig(t *testing.T) {
	loadConfig()
	t.Log(cfg.Address)
}
