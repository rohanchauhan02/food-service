package config

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

// ImmutableConfigInterface is an interface represents methods that can be used to get configuration values.
type ImmutableConfigInterface interface {
	GetPort() string
	GetDatabaseHost() string
	GetDatabasePort() string
	GetDatabaseUser() string
	GetDatabasePassword() string
	GetDatabaseName() string
	GetJWTSecret() string
}

// im is a struct that implements ImmutableConfigInterface.
type im struct {
	Port         string `mapstructure:"PORT"`
	DatabaseHost string `mapstructure:"DATABASE_HOST"`
	DatabasePort string `mapstructure:"DATABASE_PORT"`
	DatabaseUser string `mapstructure:"DATABASE_USER"`
	DatabasePass string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
}

// GetPort returns the port.
func (i *im) GetPort() string {
	return i.Port
}

// GetDatabaseHost returns the database host.
func (i *im) GetDatabaseHost() string {
	return i.DatabaseHost
}

// GetDatabasePort returns the database port.
func (i *im) GetDatabasePort() string {
	return i.DatabasePort
}

// GetDatabaseUser returns the database user.
func (i *im) GetDatabaseUser() string {
	return i.DatabaseUser
}

// GetDatabasePassword returns the database password.
func (i *im) GetDatabasePassword() string {
	return i.DatabasePass
}

// GetDatabaseName returns the database name.
func (i *im) GetDatabaseName() string {
	return i.DatabaseName
}

func (i *im) GetJWTSecret() string {
	return i.JWTSecret
}

// imOnce is a struct that implements sync.Once.
// myEnv is a map that stores the environment variables.
// immutable is a struct that implements ImmutableConfigInterface.
var (
	imOnce    sync.Once
	immutable im
)

// NewImmutableConfig is a factory that returns its config implementation.
func NewImmutableConfig() ImmutableConfigInterface {
	imOnce.Do(func() {
		v := viper.New()
		appEnv, exists := os.LookupEnv("APP_ENV")
		if exists {
			if appEnv == "development" {
				v.SetConfigName("app.congig.dev")
			} else {
				v.SetConfigName("app.config.prod")
			}
		} else {
			v.SetConfigName("app.config.local")
		}
		v.AddConfigPath(".")
		// v.SetEnvPrefix("GO_")
		v.AutomaticEnv()
		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
		v.Unmarshal(&immutable)
	})
	return &immutable
}
