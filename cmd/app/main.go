package main

import (
	"context"
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

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	groupRepository := repository.NewGroupRepository(db)
	groupService := service.NewGroupService(groupRepository, userService, taskService)
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
				groups.Get("/:gid", groupHandler.GetByGroupID)
				groups.Get("/:gid/users", groupHandler.GetUsersByGroupID)
				groups.Get("/:gid/posts", groupHandler.GetPostsByGroupID)
				groups.Get("/:gid/tasks", groupHandler.GetTasksByGroupID)
				groups.Put("/:gid", groupHandler.UpdateGroup)
				groups.Post("/:gid/tasks/:tid/discussion/start", groupHandler.StartDiscussion)
			}

			posts := api.Group("/posts", guardMiddleware.RequireAuthorization, guardMiddleware.AccessOnly(domain.UserRoleUser))
			{
				posts.Post("/", postHandler.CreatePost)
				posts.Get("/:pid/comments", postHandler.GetCommentsByPostID)
				posts.Put("/:pid", postHandler.UpdatePost)
				posts.Delete("/:pid", postHandler.DeletePost)
			}

			comments := api.Group("/comments", guardMiddleware.RequireAuthorization)
			{
				comments.Post("/", commentHandler.CreateComment)
				comments.Get("/:cid", commentHandler.GetByCommentID)
				comments.Put("/:cid", commentHandler.UpdateComment)
				comments.Delete("/:cid", commentHandler.DeleteComment)
			}

			tasks := api.Group("/tasks", guardMiddleware.RequireAuthorization, guardMiddleware.AccessOnly(domain.UserRoleUser))
			{
				tasks.Post("/", taskHandler.CreateTask)
				tasks.Get("/:tid/topics", taskHandler.GetTopicsByTaskID)
				tasks.Put("/:tid", taskHandler.UpdateTask)
				tasks.Delete("/:tid", taskHandler.DeleteTask)
			}

		}
	}

	log.Println(app.Listen(":8080"))
}
