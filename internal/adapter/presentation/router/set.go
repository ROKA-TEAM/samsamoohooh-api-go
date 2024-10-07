package router

import "samsamoohooh-go-api/internal/adapter/presentation/handler"

type HandlerSet struct {
	UserHandler  *handler.UserHandler
	GroupHandler *handler.GroupHandler
}
