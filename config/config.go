package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	LogLevel             string        `mapstructure:"LOG_LEVEL"`
	DBConnString         string        `mapstructure:"DB_CONN_STRING"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	RedisAddress         string        `mapstructure:"REDIS_CONN_STRING"`
	RedisPassword        string        `mapstructure:"REDIS_PW"`
	RedirectURL          string        `mapstructure:"REDIRECT_URL"`
}

// Load reads configuration from file or environment variables.
func Load(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Viper couldn't read in the config file. %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Viper could not unmarshal the configuration. %v", err)
	}
	return
}
