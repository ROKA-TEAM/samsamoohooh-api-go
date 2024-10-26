package presenter

type AuthRefreshResponse struct {
	AccessToken string `json:"accessToken"`
}

func NewAuthRefreshResponse(accessToken string) *AuthRefreshResponse {
	return &AuthRefreshResponse{
		AccessToken: accessToken,
	}
}
