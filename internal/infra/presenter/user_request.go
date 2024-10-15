package presenter

import "samsamoohooh-go-api/internal/domain"

type UserCreateRequest struct {
	Name       string `json:"name"  validate:"gte=1,lte=15"`
	Resolution string `json:"resolution" validate:"gte=0,lte=15"`
	Role       string `json:"role" validate:"oneof=ADMIN MANAGER MEMBER"`
}

func (r UserCreateRequest) ToDomain() *domain.User {
	return &domain.User{
		Name:       r.Name,
		Resolution: r.Resolution,
		Role:       domain.UserRoleType(r.Role),
	}
}

type UserUpdateRequest struct {
	Name       string `json:"name"  validate:"gte=1,lte=15,omitempty"`
	Resolution string `json:"resolution" validate:"gte=0,lte=15,omitempty"`
}

func (r UserUpdateRequest) ToDomain() *domain.User {
	return &domain.User{
		Name:       r.Name,
		Resolution: r.Resolution,
	}
}
