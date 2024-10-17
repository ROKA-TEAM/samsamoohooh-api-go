package google

import "samsamoohooh-go-api/internal/domain"

type exchangeResponseBody struct {
	Sub        string `json:"sub"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
}

func (r exchangeResponseBody) toDomain() *domain.OauthPayload {
	return &domain.OauthPayload{
		Sub:  r.Sub,
		Name: r.Name,
	}
}
