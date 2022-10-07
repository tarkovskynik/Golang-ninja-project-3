package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type Config struct {
	MongoDB MongoDB `mapstructure:"mongodb"`
	Server  Server  `mapstructure:"server"`
}

type MongoDB struct {
	URI      string
	Username string
	Password string
	Name     string
}

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func New(folder, filename string) (*Config, error) {
	cfg := &Config{}

	viper.AddConfigPath(folder)
	viper.SetConfigName(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("mongodb", cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	return cfg, nil
}
