package repository

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/repository/database"
	enttask "samsamoohooh-go-api/internal/application/repository/database/ent/task"
	"samsamoohooh-go-api/internal/application/repository/database/utils"
)

var _ domain.TaskRepository = (*TaskRepository)(nil)

type TaskRepository struct {
	database *database.Database
}

func NewTaskRepository(database *database.Database) *TaskRepository {
	return &TaskRepository{database: database}
}

func (r *TaskRepository) CreateTask(ctx context.Context, groupID int, task *domain.Task) (*domain.Task, error) {
	createdTask, err := r.database.Task.
		Create().
		SetDeadline(task.Deadline).
		SetRange(task.Range).
		SetGroupID(groupID).
		Save(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainTask(createdTask), nil
}
func (r *TaskRepository) GetTasks(ctx context.Context, offset, limit int) ([]*domain.Task, error) {
	listTask, err := r.database.Task.
		Query().
		Offset(offset).
		Limit(limit).
		All(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainTasks(listTask), nil
}
func (r *TaskRepository) GetByTaskID(ctx context.Context, id int) (*domain.Task, error) {
	gotTask, err := r.database.Task.
		Get(ctx, id)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainTask(gotTask), nil
}

func (r *TaskRepository) GetTopicsByTaskID(ctx context.Context, id, offset, limit int) ([]*domain.Topic, error) {
	listTopics, err := r.database.Task.
		Query().
		Where(enttask.IDEQ(id)).
		Offset(offset).
		Limit(limit).
		QueryTopics().
		All(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainTopics(listTopics), nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, id int, task *domain.Task) (*domain.Task, error) {
	updateBuilder := r.database.Task.
		UpdateOneID(id)

	if !task.Deadline.IsZero() {
		updateBuilder.SetDeadline(task.Deadline)
	}

	if task.Range != 0 {
		updateBuilder.SetRange(task.Range)
	}

	updatedTask, err := updateBuilder.Save(ctx)
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainTask(updatedTask), nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id int) error {
	err := r.database.Task.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return utils.Wrap(err)
	}

	return nil
}

func (r *TaskRepository) GetTopicsLenByTaskID(ctx context.Context, id int) (int, error) {
	cnt, err := r.database.Task.
		Query().
		Where(enttask.IDEQ(id)).
		QueryTopics().
		Count(ctx)

	if err != nil {
		return 0, utils.Wrap(err)
	}

	return cnt, nil
}
