package config

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func setup(env string) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName(env)
	viper.AddConfigPath("config")
}

// Read returns the configuration values,
//		based on the configuration files and environment variables.
func Read(env string) (Config, error) {

	var config Config

	setup(env)
	err := viper.ReadInConfig()
	if err != nil {
		return config, errors.WithStack(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, errors.WithStack(err)
	}

	return config, nil
}
