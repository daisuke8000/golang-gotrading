package main

import (
	"fmt"
	"gotrading_application/112_sma/bitflyer"
	"gotrading_application/112_sma/config"
	"gotrading_application/112_sma/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
