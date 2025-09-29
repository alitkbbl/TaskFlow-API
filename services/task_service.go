package services

import (
	"TaskFlowAPI/models"
	"TaskFlowAPI/repository"
	"context"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(ctx context.Context, task *models.Task) error {
	return s.taskRepo.Create(ctx, task)
}

func (s *TaskService) GetTask(ctx context.Context, id string) (*models.Task, error) {
	return s.taskRepo.FindByID(ctx, id)
}

func (s *TaskService) ListTasks(ctx context.Context, filters map[string]string) ([]models.Task, error) {
	return s.taskRepo.List(ctx, filters)
}

func (s *TaskService) UpdateTask(ctx context.Context, task *models.Task) error {
	return s.taskRepo.Update(ctx, task)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.taskRepo.Delete(ctx, id)
}
