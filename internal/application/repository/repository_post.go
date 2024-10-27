package repository

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/application/repository/utils"
	"samsamoohooh-go-api/internal/infra/storage/mysql"
	entpost "samsamoohooh-go-api/internal/infra/storage/mysql/ent/post"
)

var _ port.PostRepository = (*PostRepository)(nil)

type PostRepository struct {
	database *mysql.MySQL
}

func NewPostRepository(database *mysql.MySQL) *PostRepository {
	return &PostRepository{database: database}
}

func (r PostRepository) CreatePost(ctx context.Context, groupID int, post *domain.Post) (*domain.Post, error) {
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

func (r PostRepository) GetPosts(ctx context.Context, offset, limit int) ([]*domain.Post, error) {
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

func (r PostRepository) GetByPostID(ctx context.Context, id int) (*domain.Post, error) {
	gotPost, err := r.database.Post.
		Get(ctx, id)
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainPost(gotPost), nil
}

func (r PostRepository) GetCommentsByPostID(ctx context.Context, id, offset, limit int) ([]*domain.Comment, error) {
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

func (r PostRepository) UpdatePost(ctx context.Context, id int, post *domain.Post) (*domain.Post, error) {
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

func (r PostRepository) DeletePost(ctx context.Context, id int) error {
	err := r.database.Post.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return utils.Wrap(err)
	}

	return nil
}
