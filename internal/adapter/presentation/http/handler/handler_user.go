package handler

import "samsamoohooh-go-api/internal/core/port"

type UserHandler struct {
	userBusiness port.UserBusiness
}

func NewUserHandler(userBusiness port.UserBusiness) *UserHandler {
	return &UserHandler{
		userBusiness: userBusiness,
	}
}
