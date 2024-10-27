package presenter

import "samsamoohooh-go-api/internal/application/domain"

type UserGetGroupsByMe struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type UserUpdateByMeRequest struct {
	Name       string `json:"name"  validate:"min=0,max=15,omitempty"`
	Resolution string `json:"resolution" validate:"min=0,max=22,omitempty"`
}

func (r UserUpdateByMeRequest) ToDomain() *domain.User {
	return &domain.User{
		Name:       r.Name,
		Resolution: r.Resolution,
	}
}
