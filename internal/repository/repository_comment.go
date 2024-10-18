package repository

import (
	"context"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/repository/database"
	"samsamoohooh-go-api/internal/repository/database/utils"
)

var _ domain.CommentRepository = (*CommentRepository)(nil)

type CommentRepository struct {
	database *database.Database
}

func NewCommentRepository(database *database.Database) *CommentRepository {
	return &CommentRepository{database: database}
}

func (r *CommentRepository) Create(ctx context.Context, postID int, comment *domain.Comment) (*domain.Comment, error) {
	createdComment, err := r.database.Comment.
		Create().
		SetContent(comment.Content).
		SetPostID(postID).
		Save(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainComment(createdComment), nil
}
func (r *CommentRepository) List(ctx context.Context, offset, limit int) ([]*domain.Comment, error) {
	listComment, err := r.database.Comment.
		Query().
		Offset(offset).
		Limit(limit).
		All(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainComments(listComment), nil
}
func (r *CommentRepository) GetByID(ctx context.Context, id int) (*domain.Comment, error) {
	gotComment, err := r.database.Comment.
		Get(ctx, id)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainComment(gotComment), nil
}

func (r *CommentRepository) Update(ctx context.Context, id int, comment *domain.Comment) (*domain.Comment, error) {
	updateBuilder := r.database.Comment.
		UpdateOneID(id)

	if comment.Content != "" {
		updateBuilder = updateBuilder.SetContent(comment.Content)
	}

	updatedComment, err := updateBuilder.Save(ctx)
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainComment(updatedComment), nil
}

func (r *CommentRepository) Delete(ctx context.Context, id int) error {
	err := r.database.Comment.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return utils.Wrap(err)
	}

	return nil
}
