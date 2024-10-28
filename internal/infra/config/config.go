package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Server struct {
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
			GetUserInfoURL string
		}
		Kakao struct {
			GetUserInfoURL string
		}
	}

	Logger struct {
		App struct {
			Filename string

			// DebugLevel = -1
			// InfoLevel = 0
			// WarnLevel = 1
			// ErrorLevel = 2
			// DPanicLevel = 3
			// PanicLevel = 4
			// FatalLevel = 5
			Level int

			MaxSize    int  // 최대 사이즈 (MB)
			MaxBackups int  // 최대 백업 파일 수
			MaxAge     int  // 최대 파일 유지 기간 (일)
			Compress   bool //  압축 여부
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
