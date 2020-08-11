package main

import (
	"gotrading_application/112_sma/config"
	"gotrading_application/112_sma/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
	//fmt.Println(config.Config.ApiKey)
	//fmt.Println(config.Config.ApiSecret)
}
