package router

import (
	"samsamoohooh-go-api/internal/adapter/presentation/handler"
	"samsamoohooh-go-api/internal/adapter/presentation/router/middleware"
)

type HandlerSet struct {
	UserHandler  *handler.UserHandler
	GroupHandler *handler.GroupHandler
	AuthHandler  *handler.AuthHandler
}

type MiddlewareSet struct {
	GuardMiddleware *middleware.GuardMiddleware
}
