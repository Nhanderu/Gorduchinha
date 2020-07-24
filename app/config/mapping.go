package config

import (
	"time"
)

type Config struct {
	App        AppConfig        `mapstructure:"app"`
	DB         DBConfig         `mapstructure:"db"`
	Cache      CacheConfig      `mapstructure:"cache"`
	Tracking   TrackingConfig   `mapstructure:"tracking"`
	Messaging  MessagingConfig  `mapstructure:"messaging"`
	Monitoring MonitoringConfig `mapstructure:"monitoring"`
	Server     ServerConfig     `mapstructure:"server"`
	Log        LogConfig        `mapstructure:"log"`
	HTTPClient HTTPClientConfig `mapstructure:"http-client"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"-"`
	Debug   bool   `mapstructure:"debug"`
}

type DBConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
	Name string `mapstructure:"name"`
}

type CacheConfig struct {
	Host              string        `mapstructure:"host"`
	Port              int           `mapstructure:"port"`
	User              string        `mapstructure:"user"`
	Pass              string        `mapstructure:"pass"`
	DB                int           `mapstructure:"db"`
	Prefix            string        `mapstructure:"prefix"`
	DefaultExpiration time.Duration `mapstructure:"default-expiration"`
}

type TrackingConfig struct {
}

type MessagingConfig struct {
}

type MonitoringConfig struct {
}

type ServerConfig struct {
	Port                 int              `mapstructure:"port"`
	Prefix               string           `mapstructure:"prefix"`
	MaxRequestsPerSecond float64          `mapstructure:"max-requests-per-second"`
	Auth                 ServerAuthConfig `mapstructure:"auth"`
}

type ServerAuthConfig struct {
	ClientsURLs []string `mapstructure:"clients-urls"`
}

type LogConfig struct {
	LogToFile bool   `mapstructure:"log-to-file"`
	Path      string `mapstructure:"path"`
}

type HTTPClientConfig struct {
	Timeout time.Duration `mapstructure:"timeout"`
}