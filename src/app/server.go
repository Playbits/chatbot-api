package app

import (
	"flag"

	"exampleapp.com/m/v2/config"
)

func Init()  {
	configFlag := flag.String("config", "dev", "config flag can be dev for develop or prod for production")
	prodConfigPath := flag.String("config-path", "", "config-path production config file path")
	// init service configs
	config.Init(configFlag, prodConfigPath)
	port := config.Configs.Server.Port
	r := NewRouter()
	r.Run(port)
}