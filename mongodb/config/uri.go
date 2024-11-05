package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const (
	defaultURI                  = ""
	defaultDatabase             = "test"
	defaultPassword             = ""
	defaultUser                 = ""
	defaultParamter             = ""
	defaultHost                 = "127.0.0.1:27017"
	defaultEarlyRefreshInterval = 10
)

type URI struct {
	viper                *viper.Viper
	EarlyRefreshInterval time.Duration
}

func NewURI() *URI {
	u := &URI{}
	v := viper.New()
	v.AutomaticEnv()

	v.SetDefault("MONGODB_DATABASE", defaultDatabase)
	v.SetDefault("MONGODB_URI", defaultURI)
	v.SetDefault("MONGODB_PASSWORD", defaultPassword)
	v.SetDefault("MONGODB_USER", defaultUser)
	v.SetDefault("MONGODB_PARAMETER", defaultParamter)
	v.SetDefault("MONGODB_HOST", defaultHost)
	v.SetDefault("MONGODB_EARLY_REFRESH_INTERVAL", defaultEarlyRefreshInterval)

	u.viper = v
	return u
}

func (u *URI) GetURI() (string, error) {
	uri := u.viper.GetString("MONGODB_URI")
	buf := strings.Split(uri, "//")

	if cString, err := connstring.Parse(uri); err == nil {
		if len(cString.Username) == 0 && len(cString.Password) == 0 {
			if len(buf) > 1 && u.viper.GetString("MONGODB_USER") != "" && u.viper.GetString("MONGODB_PASSWORD") != "" {
				uri = fmt.Sprintf("%s//%s:%s@%s", buf[0], u.viper.GetString("MONGODB_USER"), u.viper.GetString("MONGODB_PASSWORD"), buf[1])
			}
		}
	}

	if uri == defaultURI {
		host := u.viper.GetString("MONGODB_HOST")
		parameter := u.viper.GetString("MONGODB_PARAMETER")
		user := u.viper.GetString("MONGODB_USER")
		password := u.viper.GetString("MONGODB_PASSWORD")

		if parameter != "" && parameter[0] != '?' {
			parameter = "?" + parameter
		}

		mongoDBLogin := user
		if password != "" && mongoDBLogin != "" {
			mongoDBLogin = mongoDBLogin + ":" + password
		}

		if mongoDBLogin != "" {
			mongoDBLogin = mongoDBLogin + "@"
		}

		database := u.viper.GetString("MONGODB_DATABASE")
		uri = "mongodb://" + mongoDBLogin + host + "/" + database + parameter

	}

	return uri, nil
}
