package repository

import (
	"context"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/repository/database"
	entpost "samsamoohooh-go-api/internal/repository/database/ent/post"
	"samsamoohooh-go-api/internal/repository/database/utils"
)

var _ domain.PostRepository = (*PostRepository)(nil)

type PostRepository struct {
	database *database.Database
}

func NewPostRepository(database *database.Database) *PostRepository {
	return &PostRepository{database: database}
}

func (r PostRepository) Create(ctx context.Context, groupID int, post *domain.Post) (*domain.Post, error) {
	createdPost, err := r.database.Post.
		Create().
		SetTitle(post.Title).
		SetContent(post.Content).
		SetGroupID(groupID).
		Save(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainPost(createdPost), nil
}

func (r PostRepository) List(ctx context.Context, offset, limit int) ([]*domain.Post, error) {
	listPost, err := r.database.Post.
		Query().
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainPosts(listPost), nil
}

func (r PostRepository) GetByID(ctx context.Context, id int) (*domain.Post, error) {
	gotPost, err := r.database.Post.
		Get(ctx, id)
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainPost(gotPost), nil
}

func (r PostRepository) GetCommentsByID(ctx context.Context, id, offset, limit int) ([]*domain.Comment, error) {
	listPost, err := r.database.Post.
		Query().
		Where(entpost.IDEQ(id)).
		QueryComments().
		Offset(offset).
		Limit(limit).
		All(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainComments(listPost), nil
}

func (r PostRepository) Update(ctx context.Context, id int, post *domain.Post) (*domain.Post, error) {
	updateBuilder := r.database.Post.
		UpdateOneID(id)

	if post.Title != "" {
		updateBuilder.SetTitle(post.Title)
	}

	if post.Content != "" {
		updateBuilder.SetContent(post.Content)
	}

	updatedPost, err := updateBuilder.
		Save(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainPost(updatedPost), nil
}

func (r PostRepository) Delete(ctx context.Context, id int) error {
	err := r.database.Post.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return utils.Wrap(err)
	}

	return nil
}
