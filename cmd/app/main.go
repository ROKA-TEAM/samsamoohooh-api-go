package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/handler"
	"samsamoohooh-go-api/internal/infra/catcher"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/middleware"
	"samsamoohooh-go-api/internal/infra/token"
	"samsamoohooh-go-api/internal/repository"
	"samsamoohooh-go-api/internal/repository/database"
	"samsamoohooh-go-api/internal/service"
)

func main() {
	cfg, err := config.NewConfig(".env.toml")
	if err != nil {
		log.Panicf("failed to load config: %v\n", err)
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Panicf("failed to connect to database: %v\n", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Panicf("failed to close database connection: %v\n", err)
		}
	}()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	var tokenService domain.TokenService = token.NewJWTService(cfg)
	tokenMiddleware := middleware.NewTokenMiddleware(tokenService)

	app := fiber.New(fiber.Config{
		ErrorHandler: catcher.ErrorHandler,
	})
	v1 := app.Group("v1")
	{
		api := v1.Group("/api")
		{
			useMiddleware := api.Group("", tokenMiddleware.Authorization)
			{
				userHandler.Route(useMiddleware)
			}
		}
	}

	log.Fatal(app.Listen(":8080"))

}
