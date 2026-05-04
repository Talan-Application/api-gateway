package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Services ServicesConfig `mapstructure:"services"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
}

type ServerConfig struct {
	HTTPServer HTTPServerConfig `mapstructure:"httpserver"`
}

type HTTPServerConfig struct {
	Port           int           `mapstructure:"port"`
	ReadTimeout    time.Duration `mapstructure:"readtimeout"`
	WriteTimeout   time.Duration `mapstructure:"writetimeout"`
	IdleTimeout    time.Duration `mapstructure:"idletimeout"`
	MaxHeaderBytes int           `mapstructure:"maxheaderbytes"`
}

type ServicesConfig struct {
	Auth ServiceConfig `mapstructure:"auth"`
}

type ServiceConfig struct {
	Address string `mapstructure:"address"`
}

func Load() (*Config, error) {
	_ = godotenv.Load(".env")

	raw, err := os.ReadFile("./config/config.yml")
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	expanded := os.ExpandEnv(string(raw))

	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadConfig(strings.NewReader(expanded)); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}
