package presenter

type AuthRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthValidationRequest struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
