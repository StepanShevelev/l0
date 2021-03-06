package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port    string `mapstructure:"port"`
	Url     string `mapstructure:"url"`
	Cluster string `mapstructure:"cluster"`
	Client  string `mapstructure:"client"`
	Subj    string `mapstructure:"subj"`

	DB struct {
		Username string `mapstructure:"username"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
	}
}

func New() *Config {
	return &Config{}
}

func (c *Config) Load(path string, name string, _type string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(_type)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read config error: %w", err)
	}

	err = viper.Unmarshal(c)

	if err != nil {
		return fmt.Errorf("unmarshalling config error: %w", err)
	}
	return nil
}
