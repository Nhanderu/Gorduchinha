package app

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func ReadConfig(env string) (Config, error) {

	viper.AddConfigPath("static/config")
	viper.SetConfigName(env)

	viper.AutomaticEnv()
	viper.BindEnv("server.port", "PORT")
	viper.BindEnv("db.url", "DATABASE_URL")
	viper.BindEnv("cache.url", "HEROKU_REDIS_AQUA_TLS_URL")
	viper.SetEnvKeyReplacer(strings.NewReplacer(
		".", "_",
		"-", "_",
	))

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, errors.WithStack(err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, errors.WithStack(err)
	}

	return config, nil
}

type Config struct {
	App        AppConfig        `mapstructure:"app"`
	DB         DBConfig         `mapstructure:"db"`
	Cache      CacheConfig      `mapstructure:"cache"`
	Server     ServerConfig     `mapstructure:"server"`
	HTTPClient HTTPClientConfig `mapstructure:"http-client"`
}

type AppConfig struct {
	Name  string `mapstructure:"name"`
	Debug bool   `mapstructure:"debug"`
}

type DBConfig struct {
	URL string `mapstructure:"url"`
}

type CacheConfig struct {
	URL               string        `mapstructure:"url"`
	DB                int           `mapstructure:"db"`
	Prefix            string        `mapstructure:"prefix"`
	DefaultExpiration time.Duration `mapstructure:"default-expiration"`
}

type ServerConfig struct {
	Port      int                   `mapstructure:"port"`
	Prefix    string                `mapstructure:"prefix"`
	Auth      ServerAuthConfig      `mapstructure:"auth"`
	RateLimit ServerRateLimitConfig `mapstructure:"rate-limit"`
	RouteKeys ServerRouteKeysConfig `mapstructure:"route-keys"`
}

type ServerAuthConfig struct {
	ClientsURLs []string `mapstructure:"clients-urls"`
}

type ServerRateLimitConfig struct {
	Period time.Duration `mapstructure:"period"`
	Limit  int64         `mapstructure:"limit"`
}

type ServerRouteKeysConfig struct {
	UpdateTrophies string `mapstructure:"update-trophies"`
}

type HTTPClientConfig struct {
	Timeout time.Duration `mapstructure:"timeout"`
}
