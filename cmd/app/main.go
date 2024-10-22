package main

import (
	"context"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"log"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/handler"
	"samsamoohooh-go-api/internal/application/repository"
	"samsamoohooh-go-api/internal/application/repository/database"
	"samsamoohooh-go-api/internal/application/service"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/middleware/guard"
	"samsamoohooh-go-api/internal/infra/validator"
	"samsamoohooh-go-api/pkg/box"
	"samsamoohooh-go-api/pkg/oauth/google"
	"samsamoohooh-go-api/pkg/oauth/kakao"
	"samsamoohooh-go-api/pkg/redis"
	"samsamoohooh-go-api/pkg/token/jwt"

	"github.com/gofiber/fiber/v3"
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

	// TODO: set time out
	rds, err := redis.NewRedis(context.Background(), cfg)
	if err != nil {
		log.Panicf("failed to connect to redis: %v\n", err)
	}
	defer func(rds *redis.Redis) {
		err := rds.Close()
		if err != nil {
			log.Panicf("failed to close redis connection: %v\n", err)
		}
	}(rds)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	topicRepository := repository.NewTopicRepository(db)
	topicService := service.NewTopicService(topicRepository)
	topicHandler := handler.NewTopicHandler(topicService)

	groupRepository := repository.NewGroupRepository(db)
	groupService := service.NewGroupService(groupRepository, rds, userService, taskService)
	groupHandler := handler.NewGroupHandler(groupService)

	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService)

	jwtService := jwt.NewService(cfg)
	kakaoOauthService := kakao.NewService(jwtService, userService, cfg)
	googleOauthService := google.NewService(jwtService, userService, cfg)
	guardMiddleware := guard.NewMiddleware(jwtService, userService)

	authHandler := handler.NewAuthHandler(kakaoOauthService, googleOauthService, jwtService)

	app := fiber.New(fiber.Config{
		ErrorHandler:    box.GetFiberErrorHandler(),
		StructValidator: validator.New(),
	})

	app.Use(recoverer.New())
	app.Use(helmet.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	v1 := app.Group("v1")
	{
		api := v1.Group("/api")
		{

			auth := api.Group("/auth")
			{
				authHandler.Route(auth)
			}

			users := api.Group("/users")
			{
				userHandler.Route(users, guardMiddleware)
			}

			groups := api.Group("/groups", guardMiddleware.RequireAuthorization, guardMiddleware.AccessOnly(domain.UserRoleUser))
			{
				groupHandler.Route(groups)
			}

			posts := api.Group("/posts", guardMiddleware.RequireAuthorization, guardMiddleware.AccessOnly(domain.UserRoleUser))
			{
				postHandler.Route(posts)
			}

			comments := api.Group("/comments", guardMiddleware.RequireAuthorization, guardMiddleware.AccessOnly(domain.UserRoleUser))
			{
				commentHandler.Route(comments)
			}

			tasks := api.Group("/tasks", guardMiddleware.RequireAuthorization, guardMiddleware.AccessOnly(domain.UserRoleUser))
			{
				taskHandler.Route(tasks)
			}

			topics := api.Group("/topics", guardMiddleware.RequireAuthorization, guardMiddleware.AccessOnly(domain.UserRoleUser))
			{
				topicHandler.Route(topics)
			}
		}
	}

	var listenConfig fiber.ListenConfig
	if !cfg.HTTP.Development {
		listenConfig = fiber.ListenConfig{
			CertFile:    cfg.HTTP.TLS.CertFilePath,
			CertKeyFile: cfg.HTTP.TLS.KeyFilePath,
		}
	}

	log.Println(app.Listen(":8080", listenConfig))
}
