package kakao

import (
	"samsamoohooh-go-api/internal/domain"
	"strconv"
	"time"
)

type exchangeRespBody struct {
	ID          int       `json:"id"`
	ConnectedAt time.Time `json:"connected_at"`
	Properties  struct {
		Nickname string `json:"nickname"`
	} `json:"properties"`
	KakaoAccount struct {
		ProfileNicknameNeedsAgreement bool `json:"profile_nickname_needs_agreement"`
		Profile                       struct {
			Nickname          string `json:"nickname"`
			IsDefaultNickname bool   `json:"is_default_nickname"`
		} `json:"profile"`
	} `json:"kakao_account"`
}

func (r *exchangeRespBody) toDomain() *domain.OauthPayload {
	return &domain.OauthPayload{
		Sub:  strconv.Itoa(r.ID),
		Name: r.KakaoAccount.Profile.Nickname,
	}
}
