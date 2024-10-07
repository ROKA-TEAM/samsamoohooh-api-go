package main

import (
	"log"
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

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r := router.New(c, router.HandlerSet{
		UserHandler: userHandler,
	})

	if err := r.Start(); err != nil {
		log.Panicf("server 시작 중(후)에 실패했습니다: %v\n", err)
	}
}
