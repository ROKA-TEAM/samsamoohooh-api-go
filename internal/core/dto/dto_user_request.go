package dto

import (
	"samsamoohooh-go-api/internal/core/domain"
)

type UserCreateRequest struct {
	Name       string `json:"name"`
	Resolution string `json:"resolution"`
	Role       string `json:"role"`
	Sub        string `json:"sub"`
	Social     string `json:"social"`
}

func (r UserCreateRequest) ToDomain() *domain.User {
	return &domain.User{
		Name:       r.Name,
		Resolution: r.Resolution,
		Role:       domain.RoleType(r.Role),
		Sub:        r.Sub,
		Social:     domain.SocialType(r.Social),
	}
}

type UserUpdateRequest struct {
	Name       string `json:"name" validate:"gte=2,lte=12"`
	Resolution string `json:"resolution" validate:"gte=0,lte=13"`
}

func (r *UserUpdateRequest) ToDomain() *domain.User {
	return &domain.User{
		Name:       r.Name,
		Resolution: r.Resolution,
	}
}
