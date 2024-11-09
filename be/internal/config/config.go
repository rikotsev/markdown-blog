package config

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type (
	ServerConfig struct {
		Address string
	}

	DatabaseConfig struct {
		Url     string
		MaxConn int
	}
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
	}
)

//go:embed config.toml
var configFile embed.FS

func InitConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("toml")

	data, err := configFile.ReadFile("config.toml")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded config.toml with: %w", err)
	}

	if err := v.ReadConfig(bytes.NewReader(data)); err != nil {
		return nil, fmt.Errorf("viper failed to read the embdded config.toml content with: %w", err)
	}

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config := Config{}
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("viper failed to parse the configuration with: %w", err)
	}

	return &config, nil
}
