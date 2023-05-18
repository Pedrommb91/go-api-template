package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App      `mapstructure:"app"`
		Log      `mapstructure:"logger"`
		API      `mapstructure:"api"`
		Database `mapstructure:"database"`
	}

	App struct {
		Name    string `env-required:"true" mapstructure:"name"    env:"APP_NAME"`
		Version string `env-required:"true" mapstructure:"version" env:"APP_VERSION"`
	}

	Log struct {
		Level string `env-required:"true" mapstructure:"log_level"   env:"LOG_LEVEL"`
	}

	API struct {
		CORSAllowOrigins []string `env-required:"true" mapstructure:"cors_allow_origins" env:"API_CORS_ALLOW_ORIGINS"`
		Address          string   `env-required:"true" mapstructure:"address" env:"API_ADDRESS"`
	}

	Database struct {
		Host     string `env-required:"true" mapstructure:"host" env:"DATABASE_HOST"`
		Port     string `env-required:"true" mapstructure:"port" env:"DATABASE_PORT"`
		User     string `env-required:"true" mapstructure:"user" env:"DATABASE_USER"`
		DBName   string `env-required:"true" mapstructure:"dbname" env:"DATABASE_DBNAME"`
		Password string `env-required:"true" mapstructure:"password" env:"DATABASE_PASSWORD"`
		SslMode  string `env-required:"true" mapstructure:"sslmode" env:"DATABASE_SSLMODE"`
		Schema   string `env-required:"true" mapstructure:"schema" env:"DATABASE_SCHEMA"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, fmt.Errorf("Config error %s", err)
	}

	return cfg, nil
}
