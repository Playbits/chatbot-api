package config

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
}
type DatabaseConfig struct {
    Host string `yaml:"host"`
    Port string `yaml:"port"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    Database string `yaml:"database"`
    Engine string `yaml:"engine"`
}
type configs struct {
    Server  ServerConfig `yaml:"server"`
    Database DatabaseConfig `yaml:"database"`
}

var Configs configs

func Init(Config, ConfigPath *string) {
	var configPath string
	if Config == nil || *Config == "dev" {
		_, b, _, _ := runtime.Caller(0)
		BasePath := filepath.Dir(b)
		configPath = BasePath + "/dev.yaml"
	} else {
		configPath = *ConfigPath
	}
	load(configPath)
}


func load(ConfigsPath string) {
	yamlFile, err := ioutil.ReadFile(ConfigsPath)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &Configs)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}