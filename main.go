package main

import (
	"github.com/songyanping/gin-simple/configs"
	"github.com/songyanping/gin-simple/routes"
)

func main() {

	configs.InitConfig()
	r := routes.NewRouter()
	r.Run(configs.Conf.Server.HttpPort)

}
