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

// TokenValidity defines the validity period and activation delay for tokens
type Duration struct {
	ValidityPeriod  customDuration // 토큰의 유효 기간
	ActivationDelay customDuration // 토큰이 유효해지기까지의 지연 시간
}

type OauthAuthorizationCodeGrantInfo struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

type Config struct {
	HTTP struct {
		Development bool
		Port        string
		TLS         struct {
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
		Google OauthAuthorizationCodeGrantInfo
		Kakao  OauthAuthorizationCodeGrantInfo
	}

	Logger struct {
		Encoding string
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
