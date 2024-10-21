package v1

type AuthValidationRequest struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

type AuthRefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}
