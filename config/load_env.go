package config

import (
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         int    `mapstructure:"POSTGRES_PORT"`
	DBUrl          string `mapstructure:"DATABASE_URL"`
	ServerPort     string `mapstructure:"API_PORT"`
	APIVersion     string `mapstructure:"API_VERSION"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
	ApiURL       string `mapstructure:"API_URL"`
	RedisUri     string `mapstructure:"REDIS_URL"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`

	StripeSecret string `mapstructure:"STRIPE_SECRET"`
	StripeKey    string `mapstructure:"STRIPE_KEY"`
	StripeApi    string `mapstructure:"STRIPE_API"`

	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`
	SmtpPort     string `mapstructure:"SMTP_PORT"`
	SmtpUser     string `mapstructure:"SMTP_USER"`
}

var (
	config      *Config
	configError error
	once        sync.Once
)

const (
	ConfigDefaultName = "app.env"
)

func LoadConfig(filepath string) (*Config, error) {
	once.Do(func() {

		viper.SetConfigFile(filepath)
		viper.AutomaticEnv()

		configError = viper.ReadInConfig()
		if configError != nil {
			log.Error("Error to read configs: ", configError)
			return
		}

		config = &Config{}
		configError = viper.Unmarshal(config)
		if configError != nil {
			log.Error("Error to unmarshal configs: ", configError)
			return
		}

		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			if in.Op == fsnotify.Write {
				err := viper.Unmarshal(config)
				if err != nil {
					log.Error("Error to unmarshal new config changes: ", err)
					return
				}
			}
		})
	})

	return config, configError
}
