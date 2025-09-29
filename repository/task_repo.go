package repository

import (
	"TaskFlowAPI/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, task *models.Task) error {
	if task.ID == "" {
		task.ID = uuid.NewString()
	}
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *TaskRepository) FindByID(ctx context.Context, id string) (*models.Task, error) {
	var task models.Task
	if err := r.db.WithContext(ctx).First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) List(ctx context.Context, filters map[string]string) ([]models.Task, error) {
	var tasks []models.Task
	q := r.db.WithContext(ctx).Model(&models.Task{})

	if status, ok := filters["status"]; ok {
		q = q.Where("status = ?", status)
	}
	if qstr, ok := filters["q"]; ok {
		q = q.Where("title LIKE ?", "%"+qstr+"%")
	}

	err := q.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) Update(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Save(task).Error
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.Task{}, "id = ?", id).Error
}
