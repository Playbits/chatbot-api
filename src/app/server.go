package app

import (
	"exampleapp.com/m/v2/config"
)

func Init()  {
	port := config.Configs.Server.Port
	r := NewRouter()
	r.Run(port)
}