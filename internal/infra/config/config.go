package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	HTTP struct {
		Port string
		TLS  struct {
			CertFilePath string
			KeyFilePath  string
		}
	}
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		Database string
	}

	Redis struct {
		Addr     string
		Password string
		DB       int
		Protocol int
	}

	Token struct {
		Issuer    string
		SecretKey string
		Duration  struct {
			Access  Duration
			Refresh Duration
		}
	}

	Oauth struct {
		Google struct {
			ClientID     string
			ClientSecret string
			RedirectURL  string
		}
		Kakao struct {
			ClientID     string
			ClientSecret string
			RedirectURL  string
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
