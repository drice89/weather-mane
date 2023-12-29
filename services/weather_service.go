package services

import (
	"log"
	"os"

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
		// TODO Save field data to database
		log.Println(item.LastDataFields)
	}
}