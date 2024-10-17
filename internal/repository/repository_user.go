package repository

import (
	"context"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/repository/database"
	entuser "samsamoohooh-go-api/internal/repository/database/ent/user"
	"samsamoohooh-go-api/internal/repository/database/utils"
)

var _ domain.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	database *database.Database
}

func NewUserRepository(database *database.Database) *UserRepository {
	return &UserRepository{database: database}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	createdUser, err := r.database.User.
		Create().
		SetName(user.Name).
		SetResolution(user.Resolution).
		SetRole(entuser.Role(user.Role)).
		SetSocial(entuser.Social(user.Social)).
		SetSocialSub(user.SocialSub).
		Save(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainUser(createdUser), nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	gotUser, err := r.database.User.
		Get(ctx, id)
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainUser(gotUser), nil
}
func (r *UserRepository) GetBySub(ctx context.Context, sub string) (*domain.User, error) {
	gotUser, err := r.database.User.
		Query().
		Where(entuser.SocialSubEQ(sub)).
		First(ctx)
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainUser(gotUser), nil
}

func (r *UserRepository) GetGroupsByID(ctx context.Context, id int, limit, offset int) ([]*domain.Group, error) {
	gotGroups, err := r.database.User.
		Query().
		Where(entuser.IDEQ(id)).
		QueryGroups().
		Offset(offset).
		Limit(limit).
		All(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainGroups(gotGroups), nil
}

func (r *UserRepository) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	gotUsers, err := r.database.User.
		Query().
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainUsers(gotUsers), nil
}

func (r *UserRepository) Update(ctx context.Context, id int, user *domain.User) (*domain.User, error) {
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
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainUser(updatedUser), nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	err := r.database.User.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return utils.Wrap(err)
	}

	return nil
}
