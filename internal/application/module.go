package application

import (
	"samsamoohooh-go-api/internal/application/handler"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/application/repository"
	"samsamoohooh-go-api/internal/application/service"

	"go.uber.org/fx"
)

var RepositoryModule = fx.Module(
	"repository-module",
	fx.Provide(
		fx.Annotate(
			repository.NewUserRepository,
			fx.As(new(port.UserRepository)),
		),

		fx.Annotate(
			repository.NewGroupRepository,
			fx.As(new(port.GroupRepository)),
		),

		fx.Annotate(
			repository.NewPostRepository,
			fx.As(new(port.PostRepository)),
		),

		fx.Annotate(
			repository.NewCommentRepository,
			fx.As(new(port.CommentRepository)),
		),

		fx.Annotate(
			repository.NewTaskRepository,
			fx.As(new(port.TaskRepository)),
		),

		fx.Annotate(
			repository.NewTopicRepository,
			fx.As(new(port.TopicRepository)),
		),
	),
)

var ServiceModule = fx.Module(
	"service-module",
	fx.Provide(
		fx.Annotate(
			service.NewUserService,
			fx.As(new(port.UserService)),
		),

		fx.Annotate(
			service.NewGroupService,
			fx.As(new(port.GroupService)),
		),

		fx.Annotate(
			service.NewPostService,
			fx.As(new(port.PostService)),
		),

		fx.Annotate(
			service.NewCommentService,
			fx.As(new(port.CommentService)),
		),

		fx.Annotate(
			service.NewTaskService,
			fx.As(new(port.TaskService)),
		),

		fx.Annotate(
			service.NewTopicService,
			fx.As(new(port.TopicService)),
		),
	),
)

var HandlerModule = fx.Module(
	"handler-module",
	fx.Provide(
		handler.NewErrorHandler,
		handler.NewAuthHandler,
		fx.Annotate(
			handler.NewOauthHandler,
			fx.ParamTags(`name:"kakao"`, `name:"google"`),
		),

		handler.NewUserHandler,
		handler.NewGroupHandler,
		handler.NewPostHandler,
		handler.NewCommentHandler,
		handler.NewTaskHandler,
		handler.NewTopicHandler,
	),
)
