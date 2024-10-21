package main

import (
	"context"
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
	// TODO: set time out
	if db.AutoMigration(context.Background()) != nil {
		log.Panicf("failed to auto migration: %v\n", err)
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

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)

	groupRepository := repository.NewGroupRepository(db)
	groupService := service.NewGroupService(groupRepository, userService, taskService)
	groupHandler := handler.NewGroupHandler(groupService)

	jwtService := jwt.NewService(cfg)
	kakaoOauthService := kakao.NewService(jwtService, userService, cfg)
	googleOauthService := google.NewService(jwtService, userService, cfg)
	guardMiddleware := guard.NewMiddleware(jwtService, userService)

	authHandler := handler.NewAuthHandler(kakaoOauthService, googleOauthService, jwtService)

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

			groups := api.Group("/groups", guardMiddleware.RequireAuthorization, guardMiddleware.AccessOnly(domain.UserRoleUser))
			{
				groups.Post("/", groupHandler.CreateGroup)
				groups.Get("/:gid", groupHandler.GetByGroupID, guardMiddleware.CheckGroupAccess)
				groups.Get("/:gid/users", groupHandler.GetUsersByGroupID, guardMiddleware.CheckGroupAccess)
				groups.Get("/:gid/posts", groupHandler.GetPostsByGroupID, guardMiddleware.CheckGroupAccess)
				groups.Get("/:gid/tasks", groupHandler.GetTasksByGroupID, guardMiddleware.CheckGroupAccess)
				groups.Put("/:gid", groupHandler.UpdateGroup, guardMiddleware.CheckGroupAccess)
				groups.Post("/:gid/tasks/:tid/discussion/start", groupHandler.StartDiscussion, guardMiddleware.CheckGroupAccess)
			}
		}
	}

	log.Println(app.Listen(":8080"))
}
