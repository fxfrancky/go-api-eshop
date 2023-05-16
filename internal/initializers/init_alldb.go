package initializers

import (
	"github.com/fxfrancky/go-api-eshop/config"
	"gorm.io/gorm"
)

func LoadDatabases(conf *config.Config) *gorm.DB {

	db := ConnectDB(conf)
	ConnectRedis(conf)

	return db
}
