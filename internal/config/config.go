package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Config struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		Database string
	}

	Token struct {
		Key      string
		Duration struct {
			Access  string
			Refresh string
			Temp    string
		}
	}

	HTTP struct {
		Port string
	}

	Oauth struct {
		Google struct {
			ClientID     string
			ClientSecret string
			CallbackURL  string
			UserInfoURL  string
		}
	}
}

func New(path string) (*Config, error) {
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
