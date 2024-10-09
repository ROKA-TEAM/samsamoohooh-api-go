package dto

import "samsamoohooh-go-api/internal/core/domain"

type AuthMoreInfoRequest struct {
	Name       string `json:"name" validate:"gte=2,lte=12"`
	Resolution string `json:"resolution" validate:"gte=0,lte=13"`
}

func (r *AuthMoreInfoRequest) ToDomain() *domain.User {
	return &domain.User{
		Name:       r.Name,
		Resolution: r.Resolution,
	}
}
