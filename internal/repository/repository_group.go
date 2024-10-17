package repository

import (
	"context"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/repository/database"
	groupent "samsamoohooh-go-api/internal/repository/database/ent/group"
	"samsamoohooh-go-api/internal/repository/database/utils"
)

var _ domain.GroupRepository = (*GroupRepository)(nil)

type GroupRepository struct {
	database *database.Database
}

func NewGroupRepository(database *database.Database) *GroupRepository {
	return &GroupRepository{database: database}
}

func (r *GroupRepository) Create(ctx context.Context, userID int, group *domain.Group) (*domain.Group, error) {
	createdGroup, err := r.database.Client.Group.
		Create().
		SetBookTitle(group.BookTitle).
		SetAuthor(group.Author).
		SetMaxPage(group.MaxPage).
		SetPublisher(group.Publisher).
		SetDescription(group.Description).
		SetBookMark(0).
		AddUserIDs(userID).
		Save(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainGroup(createdGroup), nil
}
func (r *GroupRepository) List(ctx context.Context, offset, limit int) ([]*domain.Group, error) {
	listGroups, err := r.database.Client.Group.
		Query().
		Offset(offset).
		Limit(limit).
		All(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainGroups(listGroups), nil
}

func (r *GroupRepository) GetByID(ctx context.Context, id int) (*domain.Group, error) {
	gotGroup, err := r.database.Client.
		Group.Get(ctx, id)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainGroup(gotGroup), nil
}
func (r *GroupRepository) GetUsersByID(ctx context.Context, id int, offset, limit int) ([]*domain.User, error) {
	listUser, err := r.database.Client.Group.
		Query().
		Where(groupent.ID(id)).
		QueryUsers().
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainUsers(listUser), nil

}
func (r *GroupRepository) GetPostsByID(ctx context.Context, id int, offset, limit int) ([]*domain.Post, error) {
	listPost, err := r.database.Client.Group.
		Query().Where(groupent.ID(id)).
		QueryPosts().
		Offset(offset).
		Limit(limit).
		All(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainPosts(listPost), nil
}
func (r *GroupRepository) GetTasksByID(ctx context.Context, id int, offset, limit int) ([]*domain.Task, error) {
	listTask, err := r.database.Client.Group.
		Query().
		Where(groupent.ID(id)).
		QueryTasks().
		Offset(offset).
		Limit(limit).
		All(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainTasks(listTask), nil
}
func (r *GroupRepository) Update(ctx context.Context, id int, group *domain.Group) (*domain.Group, error) {
	updateBuilder := r.database.Group.
		UpdateOneID(id)

	if group.BookTitle != "" {
		updateBuilder.SetBookTitle(group.BookTitle)
	}

	if group.Author != "" {
		updateBuilder.SetAuthor(group.Author)
	}

	if group.MaxPage != 0 {
		updateBuilder.SetMaxPage(group.MaxPage)
	}

	if group.Publisher != "" {
		updateBuilder.SetPublisher(group.Publisher)
	}

	if group.Description != "" {
		updateBuilder.SetDescription(group.Description)
	}

	if group.Bookmark != 0 {
		updateBuilder.SetBookMark(group.Bookmark)
	}

	updatedGroup, err := updateBuilder.Save(ctx)
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainGroup(updatedGroup), nil
}
func (r *GroupRepository) Delete(ctx context.Context, id int) error {
	err := r.database.Group.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return utils.Wrap(err)
	}

	return nil
}
