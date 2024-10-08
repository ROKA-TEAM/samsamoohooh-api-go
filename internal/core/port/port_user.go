package port

import (
	"context"
	"samsamoohooh-go-api/internal/core/domain"
	"samsamoohooh-go-api/internal/core/dto"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	GetByID(ctx context.Context, id uint) (*domain.User, error)
	GetBySub(ctx context.Context, sub string) (*domain.User, error)
	GetGroupsByID(ctx context.Context, id uint) ([]*domain.Group, error)
	GetAll(ctx context.Context, skip, limit int) ([]domain.User, error)
	Update(ctx context.Context, id uint, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id uint) (*domain.User, error)
}

type UserService interface {
	Create(ctx context.Context, user *dto.UserCreateRequest) (*dto.UserCreateResponse, error)
	GetByID(ctx context.Context, id uint) (*dto.UserGetByIDResponse, error)
	GetBySub(ctx context.Context, sub string) (*dto.UserGetBySubResponse, error)
	GetGroupsByID(ctx context.Context, id uint) ([]*dto.UserGetGroupsByIDResponse, error)
	GetAll(ctx context.Context, skip, limit int) ([]*dto.UserGetAllResponse, error)
	Update(ctx context.Context, id uint, user *dto.UserUpdateRequest) (*dto.UserUpdateResponse, error)
	Delete(ctx context.Context, id uint) (*dto.UserDeleteResponse, error)
}
