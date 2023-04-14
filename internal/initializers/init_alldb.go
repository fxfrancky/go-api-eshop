package initializers

import (
	"log"

	"github.com/fxfrancky/go-api-eshop/config"
)

func LoadDatabases(path string) {
	conf, err := config.LoadConfig(path)
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}

	ConnectRedis(&conf)
	ConnectDB(&conf)
}
