package initializers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fxfrancky/go-api-eshop/config"
	"github.com/fxfrancky/go-api-eshop/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(config *config.Config) *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 10, // Slow SQL threshold
			LogLevel:                  logger.Info,           // Log level
			IgnoreRecordNotFoundError: false,                 // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                  // Disable color
		},
	)

	log.Println(" ************************ The DBName from config : ", config.DBName)

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	err = AutoMigrate(DB)
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	log.Println("🚀 Connected Successfully to the PostgreSQL Database")
	return DB
}

// Migrate the database changes with Gorm
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		&models.PaymentResult{},
		&models.ShippingAddress{},
		&models.Review{},
	)
	if err != nil {
		return err
	}
	return nil
}
