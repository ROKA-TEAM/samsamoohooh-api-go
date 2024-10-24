package config

import "time"

type customDuration time.Duration

func (d *customDuration) UnmarshalText(b []byte) error {
	duration, err := time.ParseDuration(string(b))
	if err != nil {
		return err
	}

	*d = customDuration(duration)
	return nil
}

type Duration struct {
	ValidityPeriod  customDuration // 토큰의 유효 기간
	ActivationDelay customDuration // 토큰이 유효해지기까지의 지연 시간
}
