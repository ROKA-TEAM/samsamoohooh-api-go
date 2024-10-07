package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Database struct {
		User            string
		Password        string
		Host            string
		Port            string
		Database        string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifeTime string
	}

	Token struct {
		Key             string
		AccessDuration  string
		RefreshDuration string
	}

	HTTP struct {
		Port string
	}
}

func New(path string) (*Config, error) {
	// ? viper 사용해서 구현하는 것도 좋을듯 하다. (https://github.com/spf13/viper)
	config := new(Config)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	err = toml.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return config, nil
}
