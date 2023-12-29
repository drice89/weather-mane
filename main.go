package main

import (
	"log"
	"wm/config"
	"wm/services"
)

func main() {
	log.Println("INFO: Initializing program")

	config.InitConfig()
	// TODO loop get weather
	services.GetWeather()
}
