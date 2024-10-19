package main

import (
	"context"
	"log"
	"samsamoohooh-go-api/internal/handler"
	"samsamoohooh-go-api/internal/infra/catcher"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/logger"
	"samsamoohooh-go-api/internal/infra/middleware"
	"samsamoohooh-go-api/internal/infra/oauth/google"
	"samsamoohooh-go-api/internal/infra/oauth/kakao"
	"samsamoohooh-go-api/internal/infra/token"
	"samsamoohooh-go-api/internal/repository"
	"samsamoohooh-go-api/internal/repository/database"
	"samsamoohooh-go-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.NewConfig(".env.toml")
	if err != nil {
		log.Panicf("failed to load config: %v\n", err)
	}

	err = logger.Initialize(cfg)
	if err != nil {
		log.Panicf("failed to initialize logger: %v\n", err)
	}
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Panicf("failed to sync logger: %v\n", err)
		}
	}()

	logger.Get().Debug("success logger initialized")

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

	logger.Get().Debug("success connect to database")

	if err := db.AutoMigration(context.Background()); err != nil {
		log.Panicf("failed to auto migrate: %v\n", err)
	}

	logger.Get().Debug("success auto migrate")

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	groupRepository := repository.NewGroupRepository(db)
	groupService := service.NewGroupService(groupRepository)
	groupHandler := handler.NewGroupHandler(groupService)

	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService)

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	topicRepository := repository.NewTopicRepository(db)
	topicService := service.NewTopicService(topicRepository)
	topicHandler := handler.NewTopicHandler(topicService)

	jwtService := token.NewJWTService(cfg)
	guardMiddleware := middleware.NewGuardMiddleware(jwtService)

	oauthGoogleService := google.NewOauthGoogleService(cfg, userService, jwtService)
	oauthKakaoService := kakao.NewOauthKakaoService(cfg, userService, jwtService)
	authHandler := handler.NewAuthHandler(oauthGoogleService, oauthKakaoService, jwtService)

	logger.Get().Debug("success dependency injection")

	app := fiber.New(fiber.Config{
		ErrorHandler: catcher.ErrorHandler,
	})

	v1 := app.Group("/v1")
	{
		api := v1.Group("/api")
		{
			auth := api.Group("/auth")
			{
				authHandler.Route(auth)
			}

			users := api.Group("/users", guardMiddleware.RequireAuthorization)
			{
				userHandler.Route(users, guardMiddleware)
			}

			groups := api.Group("/groups", guardMiddleware.RequireAuthorization)
			{
				groupHandler.Route(groups, guardMiddleware)
			}

			posts := api.Group("/posts", guardMiddleware.RequireAuthorization)
			{
				postHandler.Route(posts, guardMiddleware)
			}

			comments := api.Group("/comments", guardMiddleware.RequireAuthorization)
			{
				commentHandler.Route(comments, guardMiddleware)
			}

			tasks := api.Group("/tasks", guardMiddleware.RequireAuthorization)
			{
				taskHandler.Route(tasks, guardMiddleware)
			}

			topics := api.Group("/topics", guardMiddleware.RequireAuthorization)
			{
				topicHandler.Route(topics, guardMiddleware)
			}
		}
	}

	logger.Get().Debug("success route api")
	if cfg.HTTP.Development {
		log.Println(app.Listen(cfg.HTTP.Port))
	} else {
		log.Println(app.ListenTLS(cfg.HTTP.Port, cfg.HTTP.TLS.CertFilePath, cfg.HTTP.TLS.KeyFilePath))
	}

	// app.ListenTLS()
}
