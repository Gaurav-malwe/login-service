package config

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	conf *Config
)

type Config struct {
	*viper.Viper
}

func GetConfig() *Config {
	if conf == nil {
		once.Do(Load)
	}
	return conf
}

func Load() {
	// Load config from env variables
	conf = &Config{viper.New()}
	conf.AutomaticEnv()

	// set defaults for all env
	// application settings
	conf.SetDefault("server_port", "8080")
	conf.SetDefault("service_name", "traical")

	conf.LoadFromFile()
}

func (conf *Config) LoadFromFile() {
	v := conf.Viper
	configPath, err := os.Getwd()

	if err != nil {
		return
	}

	configPath += "/../config"

	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err = v.ReadInConfig()

	if err != nil {
		return
	}

	log.Printf("Loading config from file %s", v.ConfigFileUsed())

	// Log loading config file from
}
