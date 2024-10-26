package presenter

type OauthGoogleLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewOauthGoogleLoginResponse(accessToken, refreshToken string) *OauthGoogleLoginResponse {
	return &OauthGoogleLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

type OauthKakaoLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewOauthKakaoLoginResponse(accessToken, refreshToken string) *OauthKakaoLoginResponse {
	return &OauthKakaoLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
