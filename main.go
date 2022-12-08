package main

import (
	"flag"

	"exampleapp.com/m/v2/config"
	"exampleapp.com/m/v2/src/database"
)

func main() {
	loadConfig()
	go h.run()
	database.Init()
	Init()
}

func loadConfig()  {
	configFlag := flag.String("config", "dev", "config flag can be dev for develop or prod for production")
	prodConfigPath := flag.String("config-path", "", "config-path production config file path")
	config.Init(configFlag, prodConfigPath)
}