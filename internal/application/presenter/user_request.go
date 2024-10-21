package presenter

import (
	"samsamoohooh-go-api/internal/application/domain"
)

type UserCreateRequest struct {
	Name       string `json:"name"  validate:"min=1,max=15"`
	Resolution string `json:"resolution" validate:"min=0,max=15"`
	Role       string `json:"role" validate:"oneof=ADMIN GUEST"`
	Social     string `json:"social" validate:"oneof=GOOGLE KAKAO APPLE"`
}

func (r UserCreateRequest) ToDomain() *domain.User {
	return &domain.User{
		Name:       r.Name,
		Resolution: r.Resolution,
		Social:     domain.UserSocialType(r.Social),
		Role:       domain.UserRoleType(r.Role),
	}
}

type UserUpdateRequest struct {
	Name       string `json:"name"  validate:"min=1,max=15,omitempty"`
	Resolution string `json:"resolution" validate:"min=0,max=15,omitempty"`
}

func (r UserUpdateRequest) ToDomain() *domain.User {
	return &domain.User{
		Name:       r.Name,
		Resolution: r.Resolution,
	}
}

type UserUpdateMeRequest struct {
	Name       string `json:"name"  validate:"min=0,max=15,omitempty"`
	Resolution string `json:"resolution" validate:"min=0,max=22,omitempty"`
}

func (r UserUpdateMeRequest) ToDomain() *domain.User {
	return &domain.User{
		Name:       r.Name,
		Resolution: r.Resolution,
	}
}
