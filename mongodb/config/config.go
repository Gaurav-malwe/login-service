package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/event"
)

const (
	defaultTimeOut                = 20
	defaultMaxPoolSize            = 16
	defaultMaxConnIdleTime        = 30
	defaultHeartbeatInterval      = 15
	defaultMinPoolSize            = 1
	defaultServerSelectionTimeout = 10
	defaultAppName                = "test"
)

type Config struct {
	AppName                string
	URI                    *URI
	Database               string
	Timeout                time.Duration
	MaxPoolSize            uint64
	MinPoolSize            uint64
	MaxConnIdleTime        time.Duration
	HeartbeatInterval      time.Duration
	EarlyRefreshInterval   time.Duration
	ServerSelectionTimeout time.Duration
	PoolMonitor            *event.PoolMonitor
	CommandMonitor         *event.CommandMonitor
}

func NewConfigFromEnv() *Config {
	v := viper.New()
	LoadFromFile(v)

	u := NewURI()
	uri, err := u.GetURI()

	if err != nil {
		log.Fatalf("config NewConfigFromEnv u.GetURI error: %v", (err))
		return nil
	}

	v.SetDefault("APP_NAME", defaultAppName)
	appName := v.GetString("APP_NAME")

	v.SetDefault("MONGODB_DATABASE", defaultDatabase)
	database := v.GetString("MONGODB_DATABASE")

	v.SetDefault("MONGODB_TIMEOUT", defaultTimeOut)
	timeout := v.GetDuration("MONGODB_TIMEOUT") * time.Second

	v.SetDefault("MONGODB_MIN_POOLSIZE", defaultMaxPoolSize)
	maxPoolSize := v.GetUint64("MONGODB_MIN_POOLSIZE")

	v.SetDefault("MONGODB_MIN_POOLSIZE", defaultMinPoolSize)
	minPoolSize := v.GetUint64("MONGODB_MIN_POOLSIZE")

	v.SetDefault("MONGODB_MAX_CONN_IDLE_TIME", defaultMaxConnIdleTime)
	maxConnIdleTime := v.GetDuration("MONGODB_MAX_CONN_IDLE_TIME") * time.Second

	v.SetDefault("MONGODB_HEARTBEAT_INTERVAL", defaultHeartbeatInterval)
	heartBeatInterval := v.GetDuration("MONGODB_HEARTBEAT_INTERVAL") * time.Second

	v.SetDefault("MONGODB_EARLY_REFRESH_INTERVAL", defaultEarlyRefreshInterval)
	earlyRefreshInterval := v.GetDuration("MONGODB_EARLY_REFRESH_INTERVAL") * time.Second

	v.SetDefault("MONGODB_SERVER_SELECTION_TIMEOUT", defaultServerSelectionTimeout)
	serverSelectionTimeout := v.GetDuration("MONGODB_SERVER_SELECTION_TIMEOUT") * time.Second

	// mongodb config initialized
	log.Println("mongodb config initialized", uri)

	c := &Config{
		AppName:                appName,
		URI:                    u,
		Database:               database,
		Timeout:                timeout,
		MaxConnIdleTime:        maxConnIdleTime,
		MaxPoolSize:            maxPoolSize,
		MinPoolSize:            minPoolSize,
		HeartbeatInterval:      heartBeatInterval,
		EarlyRefreshInterval:   earlyRefreshInterval,
		ServerSelectionTimeout: serverSelectionTimeout,
	}

	return c

}

func LoadFromFile(v *viper.Viper) {

	v.AutomaticEnv()

	v.AddConfigPath("../config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		return
	}

	// Log loading config file from
}
