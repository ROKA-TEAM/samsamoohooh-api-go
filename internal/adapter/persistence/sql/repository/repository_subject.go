package repository

import (
	"context"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/database"
	"samsamoohooh-go-api/internal/core/domain"
	"samsamoohooh-go-api/internal/core/port"
)

var _ port.SubjectRepository = (*SubjectRepository)(nil)

type SubjectRepository struct {
	database *database.Database
}

func NewSubjectRepository(database *database.Database) *SubjectRepository {
	return &SubjectRepository{
		database: database,
	}
}

func (r *SubjectRepository) Create(ctx context.Context, subject *domain.Subject) (*domain.Subject, error) {
	err := r.database.WithContext(ctx).Create(subject).Error
	if err != nil {
		return nil, err
	}

	return subject, nil
}
func (r *SubjectRepository) GetByID(ctx context.Context, id uint) (*domain.Subject, error) {
	subject := domain.Subject{}
	err := r.database.WithContext(ctx).First(&subject, id).Error
	if err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *SubjectRepository) GetAll(ctx context.Context, skip, limit int) ([]domain.Subject, error) {
	var subjects []domain.Subject
	err := r.database.WithContext(ctx).Limit(limit).Offset((limit - 1) * skip).Find(&subjects).Error
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (r *SubjectRepository) Update(ctx context.Context, id uint, subject *domain.Subject) (*domain.Subject, error) {
	subject.ID = id
	err := r.database.WithContext(ctx).Save(subject).Error
	if err != nil {
		return nil, err
	}

	return subject, nil
}
func (r *SubjectRepository) Delete(ctx context.Context, id uint) error {
	err := r.database.WithContext(ctx).Delete(&domain.Subject{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
