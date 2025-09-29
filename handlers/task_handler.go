package handlers

import (
	"TaskFlowAPI/models"
	"TaskFlowAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TaskHandler وظیفه هندل کردن درخواست‌های مربوط به تسک‌ها رو داره
type TaskHandler struct {
	service *services.TaskService
}

// سازنده TaskHandler
func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// CreateTask -> ایجاد تسک جدید
func (h *TaskHandler) CreateTask(c *gin.Context) {
	userID := c.GetString("userID") // گرفتن userID از JWT
	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      "todo",
		UserID:      userID,
	}

	if err := h.service.CreateTask(c.Request.Context(), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

// GetTask -> دریافت اطلاعات یک تسک با id
func (h *TaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := h.service.GetTask(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// ListTasks -> لیست همه تسک‌ها با امکان فیلتر (status, q)
func (h *TaskHandler) ListTasks(c *gin.Context) {
	filters := map[string]string{}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if q := c.Query("q"); q != "" {
		filters["q"] = q
	}

	tasks, err := h.service.ListTasks(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// UpdateTask -> ویرایش عنوان و توضیحات تسک
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("userID")

	task, err := h.service.GetTask(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	if task.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your task"})
		return
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Title = req.Title
	task.Description = req.Description

	if err := h.service.UpdateTask(c.Request.Context(), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// DeleteTask -> حذف تسک
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("userID")

	task, err := h.service.GetTask(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	if task.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your task"})
		return
	}

	if err := h.service.DeleteTask(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateStatus -> تغییر وضعیت تسک (todo, in_progress, done)
func (h *TaskHandler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("userID")

	task, err := h.service.GetTask(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	if task.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your task"})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Status = req.Status

	if err := h.service.UpdateTask(c.Request.Context(), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}
