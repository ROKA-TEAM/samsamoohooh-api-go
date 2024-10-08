package main

import (
	"log"
	"samsamoohooh-go-api/internal/adapter/auth/jwt"
	"samsamoohooh-go-api/internal/adapter/auth/oauth/google"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/database"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/repository"
	"samsamoohooh-go-api/internal/adapter/presentation/handler"
	"samsamoohooh-go-api/internal/adapter/presentation/router"
	"samsamoohooh-go-api/internal/core/service"
	"samsamoohooh-go-api/internal/infra/config"
)

func main() {

	c, err := config.New(".toml")
	if err != nil {
		log.Panicf("config 생성에 실패하였습니다: %v\n", err)
	}

	db, err := database.NewDatabase(c)
	if err != nil {
		log.Panicf("database 생성에 실패하였습니다: %v", err)
	}

	err = database.AutoMigrate(db)
	if err != nil {
		log.Panicf("migrate에 실패하였습니다: %v", err)
	}

	// jwt service
	jwtService, err := jwt.New(c)
	if err != nil {
		log.Panicf("jwtService를 불러오는데 실패했습니다: %v", err)
	}

	_ = jwtService

	// google oauth
	googleOauthService := google.New(c)

	// setting layers
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	groupRepository := repository.NewGroupRepository(db)
	groupService := service.NewGroupService(groupRepository, userRepository)
	groupHandler := handler.NewGroupHandler(groupService)

	authService := service.NewAuthService()
	authHandler := handler.NewAuthHandler(googleOauthService, userRepository, jwtService, authService)

	r := router.New(c, router.HandlerSet{
		UserHandler:  userHandler,
		GroupHandler: groupHandler,
		AuthHandler:  authHandler,
	})

	if err := r.Start(); err != nil {
		log.Panicf("server 시작 중(후)에 실패했습니다: %v\n", err)
	}
}
