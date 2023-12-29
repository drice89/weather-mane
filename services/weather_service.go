package services

import (
	"os"
	"log"

	"github.com/lrosenman/ambient"
)

func GetWeather() {
	k := ambient.NewKey(os.Getenv("AWN_APP_KEY"), os.Getenv("AWN_API_KEY"))
	resp, err := ambient.Device(k)

	if err != nil {
		log.Panicln("unable to retrieve devices")
	}

	log.Printf("%v devices found", len(resp.DeviceRecord))
	for _, item := range resp.DeviceRecord {
		log.Printf( item.Macaddress, item.Info.Name)
	}
}