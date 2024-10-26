package presenter

type OauthGoogleLoginRequest struct {
	Token string `json:"token"`
}

type OauthKakaoLoginRequest struct {
	Token string `json:"token"`
}
