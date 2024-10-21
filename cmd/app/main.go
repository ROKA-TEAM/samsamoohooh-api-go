package main

import (
	"github.com/gofiber/fiber/v3"
	"log"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/handler"
	"samsamoohooh-go-api/internal/application/repository"
	"samsamoohooh-go-api/internal/application/repository/database"
	"samsamoohooh-go-api/internal/application/service"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/middleware/guard"
	"samsamoohooh-go-api/pkg/oauth/google"
	"samsamoohooh-go-api/pkg/oauth/kakao"
	"samsamoohooh-go-api/pkg/token/jwt"
)

func main() {
	cfg, err := config.NewConfig(".env.toml")
	if err != nil {
		log.Panicf("failed to load config: %v\n", err)
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Panicf("failed to connect database: %v\n", err)
	}
	defer func(db *database.Database) {
		err := db.Close()
		if err != nil {
			log.Panicf("failed to close database connection: %v\n", err)
		}
	}(db)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	jwtService := jwt.NewService(cfg)
	kakaoOauthService := kakao.NewService(jwtService, userService, cfg)
	googleOauthService := google.NewService(jwtService, userService, cfg)

	authHandler := handler.NewAuthHandler(kakaoOauthService, googleOauthService, jwtService)

	guardMiddleware := guard.NewMiddleware(jwtService, userService)

	app := fiber.New(fiber.Config{})

	v1 := app.Group("v1")
	{
		api := v1.Group("/api")
		{

			auth := api.Group("/auth")
			{
				token := auth.Group("/token")
				{
					token.Post("/refresh", authHandler.Refresh)
					token.Post("/validation", authHandler.Validation)
				}

				auth.Get("/google", authHandler.GetLoginURLOfGoogle)
				auth.Get("/google/callback", authHandler.GoogleCallback)
				auth.Get("/kakao", authHandler.GetLoginURLOfKakao)
				auth.Get("/kakao/callback", authHandler.KaKaoCallback)
			}

			users := api.Group("/users")
			{
				me := users.Group("/me", guardMiddleware.RequireAuthorization, guardMiddleware.AccessOnly(domain.UserRoleUser))
				{
					me.Get("/", userHandler.GetByMe)
					me.Get("/groups", userHandler.GetGroupsByMe)
					me.Put("/", userHandler.UpdateMe)
					me.Delete("/", userHandler.DeleteMe)
				}
			}
		}
	}

	log.Println(app.Listen(":8080"))
}
