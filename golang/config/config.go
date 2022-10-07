package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var DEFAULT_CONFIGS = []Config{}

type Config interface {
	Name() string
}

var ConfigFile *ini.File
var Configs = make([]Config, 0)
var Environment = CurrentEnvironment()

func LoadFromFile(filename string) {
	var err error
	ConfigFile, err = ini.Load(filename)
	if err != nil {
		fmt.Println("config not found, using default config")
		ConfigFile = ini.Empty()
	}
	for _, cfg := range DEFAULT_CONFIGS {
		LoadConfig(cfg)
	}
}

func init() {
	ConfigFile = ini.Empty()
	for _, cfg := range DEFAULT_CONFIGS {
		LoadConfig(cfg)
	}
}

func LoadConfig(cfg Config) {
	sec, err := ConfigFile.GetSection(cfg.Name())
	if err == nil {
		_ = sec.MapTo(cfg)
	}
	Configs = append(Configs, cfg)
	return
}

func SaveToConfigFile(filename string) error {
	cfgFile := ini.Empty()
	for _, cfg := range Configs {
		if err := cfgFile.Section(cfg.Name()).ReflectFrom(cfg); err != nil {
			fmt.Println(err)
		}
	}
	return cfgFile.SaveTo(filename)
}
