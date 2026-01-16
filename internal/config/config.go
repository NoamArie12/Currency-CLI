package config

import (
	"github.com/spf13/viper"
	"os/user"
	"path/filepath"
)

// Config holds the application configuration.
type Config struct {
	CopyAnswer bool `mapstructure:"copy_answer"`
}

var Cfg *Config

// LoadConfig loads the configuration from file.
func LoadConfig() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	configPath := filepath.Join(usr.HomeDir, ".config", "currency-cli")

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.SetConfigType("yaml")

	viper.SetDefault("copy_answer", false)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			return err
		}
	}

	return viper.Unmarshal(&Cfg)
}
