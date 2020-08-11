package main

import (
	"fmt"
	"gotrading_application/112_sma/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
