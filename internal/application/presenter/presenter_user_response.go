package presenter

import "samsamoohooh-go-api/internal/application/domain"

type UserGetByMeResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Resolution string `json:"resolution"`
}

func NewUserGetByMeResponse(user *domain.User) *UserGetByMeResponse {
	return &UserGetByMeResponse{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
	}
}

type UserGetGroupsByMeResponse struct {
	ID          int    `json:"id"`
	BookTitle   string `json:"bookTitle"`
	Author      string `json:"author"`
	MaxPage     int    `json:"maxPage"`
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
	Bookmark    int    `json:"bookmark"`
}

func NewUserGetGroupsByMeResponse(groups []*domain.Group) []*UserGetGroupsByMeResponse {
	var listGroups []*UserGetGroupsByMeResponse
	for _, group := range groups {
		listGroups = append(listGroups, &UserGetGroupsByMeResponse{
			ID:          group.ID,
			BookTitle:   group.BookTitle,
			Author:      group.Author,
			MaxPage:     group.MaxPage,
			Publisher:   group.Publisher,
			Description: group.Description,
			Bookmark:    group.Bookmark,
		})
	}

	return listGroups
}

type UserUpdateByMeResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Resolution string `json:"resolution"`
}

func NewUserUpdateByMeResponse(user *domain.User) *UserUpdateByMeResponse {
	return &UserUpdateByMeResponse{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
	}
}
