package main

import (
	"fmt"

	"github.com/raggaer/bison/lua"
	glua "github.com/tul/gopher-lua"
)

// Config struct used for all configuration options
type Config struct {
	Address string
	DevMode bool
	Custom  map[string]interface{}
}

func loadConfig() (*Config, error) {
	configState := glua.NewState()

	defer configState.Close()
	if err := configState.DoFile("config.lua"); err != nil {
		return nil, err
	}
	configTable := configState.Get(-1)

	// Check if returned value is table
	if !lua.IsValueTable(configTable) {
		return nil, fmt.Errorf("Invalid config.lua returned data. Expected table but got %s", configTable.Type().String())
	}

	configMap := lua.TableToMap(configTable.(*glua.LTable))
	return populateConfig(configMap), nil
}

func populateConfig(m map[string]interface{}) *Config {
	dst := &Config{}
	if address, ok := m["address"].(string); ok {
		dst.Address = address
	}
	if devMode, ok := m["devMode"].(bool); ok {
		dst.DevMode = devMode
	}
	dst.Custom = m
	return dst
}
