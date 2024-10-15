package config

import (
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"
)

type customDuration time.Duration

func (d *customDuration) UnmarshalText(b []byte) error {
	duration, err := time.ParseDuration(string(b))
	if err != nil {
		return err
	}

	*d = customDuration(duration)
	return nil
}

type Config struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		Database string
	}

	Token struct {
		SecretKey string
		Duration  struct {
			Access  customDuration
			Refresh customDuration
		}
	}
}

func NewConfig(path string) (*Config, error) {
	config := new(Config)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = toml.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
