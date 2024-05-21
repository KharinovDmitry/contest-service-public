package handlers

import (
	"contest/internal/domain/enum"
	"contest/internal/domain/model"
	"contest/internal/domain/repository"
	"contest/internal/domain/service/logger"
	"contest/internal/server/dto"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AddTaskRequest struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type TaskHandler struct {
	repository repository.TaskRepository
	log        logger.Logger
}

func NewTaskHandler(repository repository.TaskRepository, log logger.Logger) *TaskHandler {
	return &TaskHandler{
		repository: repository,
		log:        log,
	}
}

// AddTask
// @Summary Добавление задачи
// @Tags Task
// @Security Bearer
// @Accept json
// @Param input body handlers.AddTaskRequest true "Информация о задаче"
// @Success 200
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /api/task [put]
func (h *TaskHandler) AddTask(c *gin.Context) {
	role := c.Request.Context().Value(enum.ContextKeyRole)
	if role != enum.Moderator && role != enum.Admin {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	var task AddTaskRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&task); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.repository.AddTask(c, task.Title, task.Text); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.log.Error(err.Error())
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

// DeleteTask
// @Summary Удаление задачи
// @Tags Task
// @Security Bearer
// @Param id path int true "Task ID"
// @Success 200
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /api/task/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	role := c.Request.Context().Value(enum.ContextKeyRole)
	if role != enum.Moderator && role != enum.Admin {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.repository.DeleteTask(c, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.log.Error(err.Error())
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

// UpdateTask
// @Summary Изменение задачи
// @Tags Task
// @Security Bearer
// @Accept json
// @Param id path int true "Task ID"
// @Param input body handlers.AddTaskRequest true "Новая информация о задаче"
// @Success 200
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /api/task/{id} [patch]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	role := c.Request.Context().Value(enum.ContextKeyRole)
	if role != enum.Moderator && role != enum.Admin {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var task dto.TaskDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&task); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.repository.UpdateTask(c, id, model.Task(task))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.log.Error(err.Error())
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

// GetTask
// @Summary Получение задачи
// @Tags Task
// @Security Bearer
// @Param id path int true "Task ID"
// @Success 200 {object} dto.TaskDTO
// @Failure 404
// @Failure 500
// @Router /api/task/{id} [get]
func (h *TaskHandler) GetTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	task, err := h.repository.FindTaskByID(c, id)
	if err != nil {
		h.log.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.TaskDTO(task))
}

// GetAllTasks
// @Summary Получение всех задач
// @Tags Task
// @Security Bearer
// @Success 200 {object} []dto.TaskDTO
// @Failure 404
// @Failure 500
// @Router /api/tasks [get]
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.repository.GetTasks(c)
	if err != nil {
		h.log.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.TasksToTasksDTO(tasks))
}
