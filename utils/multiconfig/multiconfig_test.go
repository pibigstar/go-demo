package main

import (
	"github.com/koding/multiconfig"
	"testing"
)

type Server struct {
	Name    string
	Port    int `default:"6060"`
	Enabled bool
	Users   []string
}

// supports TOML, JSON and YAML
func TestReadConfig(t *testing.T) {

	m := multiconfig.NewWithPath("config.toml")
	serverConf := new(Server)

	err := m.Load(serverConf)
	if err != nil {
		t.Error(err)
	}
	// Panic's if there is any error
	m.MustLoad(serverConf)

	t.Logf("%+v", serverConf)
}
