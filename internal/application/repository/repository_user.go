package repository

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/repository/database"
	entgroup "samsamoohooh-go-api/internal/application/repository/database/ent/group"
	entuser "samsamoohooh-go-api/internal/application/repository/database/ent/user"
	"samsamoohooh-go-api/internal/application/repository/database/utils"
)

var _ domain.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	database *database.Database
}

func NewUserRepository(database *database.Database) *UserRepository {
	return &UserRepository{database: database}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	createdUser, err := r.database.User.
		Create().
		SetName(user.Name).
		SetResolution(user.Resolution).
		SetRole(entuser.Role(user.Role)).
		SetSocial(entuser.Social(user.Social)).
		SetSocialSub(user.SocialSub).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return utils.ConvertDomainUser(createdUser), nil
}

func (r *UserRepository) GetByUserID(ctx context.Context, id int) (*domain.User, error) {
	gotUser, err := r.database.User.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return utils.ConvertDomainUser(gotUser), nil
}
func (r *UserRepository) GetByUserSub(ctx context.Context, sub string) (*domain.User, error) {
	gotUser, err := r.database.User.
		Query().
		Where(entuser.SocialSubEQ(sub)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertDomainUser(gotUser), nil
}

func (r *UserRepository) GetGroupsByUserID(ctx context.Context, id int, limit, offset int) ([]*domain.Group, error) {
	gotGroups, err := r.database.User.
		Query().
		Where(entuser.IDEQ(id)).
		QueryGroups().
		Offset(offset).
		Limit(limit).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return utils.ConvertDomainGroups(gotGroups), nil
}

func (r *UserRepository) GetUsers(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	gotUsers, err := r.database.User.
		Query().
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertDomainUsers(gotUsers), nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int, user *domain.User) (*domain.User, error) {
	updateBuilder := r.database.User.
		UpdateOneID(id)

	if user.Name != "" {
		updateBuilder.SetName(user.Name)
	}

	if user.Resolution != "" {
		updateBuilder.SetResolution(user.Resolution)
	}

	if user.Role != "" {
		updateBuilder.SetRole(entuser.Role(user.Role))
	}

	if user.Social != "" {
		updateBuilder.SetSocial(entuser.Social(user.Social))
	}

	if user.SocialSub != "" {
		updateBuilder.SetSocialSub(user.SocialSub)
	}

	updatedUser, err := updateBuilder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertDomainUser(updatedUser), nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	err := r.database.User.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) IsUserInGroup(ctx context.Context, userID, groupID int) (bool, error) {
	cnt, err := r.database.User.
		Query().
		Where(entuser.IDEQ(userID)).
		QueryGroups().
		Where(entgroup.IDEQ(groupID)).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return cnt > 0, nil
}
