package dto

type AuthGoogleCallbackResponse struct {
	TempToken    string `json:"temp_token,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
