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
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type AddTestRequest struct {
	TaskID         int    `json:"taskID,string"`
	Input          string `json:"input"`
	ExpectedResult string `json:"expectedResult"`
	Points         int    `json:"points,string"`
}

type TestHandler struct {
	repo repository.TestRepository
	log  logger.Logger
}

func NewTestHandler(testCrudService repository.TestRepository, log logger.Logger) *TestHandler {
	return &TestHandler{
		repo: testCrudService,
		log:  log,
	}
}

// AddTest
// @Summary Добавление теста
// @Tags Test
// @Security Bearer
// @Accept json
// @Param input body handlers.AddTestRequest true "Информация о тесте"
// @Success 200
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /api/test [put]
func (h *TestHandler) AddTest(c *gin.Context) {
	role := c.Request.Context().Value(enum.ContextKeyRole)
	if role != enum.Moderator && role != enum.Admin {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	var test AddTestRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&test); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		h.log.Error(err.Error())
		return
	}

	if err := h.repo.AddTest(c, test.TaskID, test.Input, test.ExpectedResult, test.Points); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.log.Error(err.Error())
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

// DeleteTest
// @Summary Удаление теста
// @Tags Test
// @Security Bearer
// @Param id path int true "Test ID"
// @Success 200
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /api/test/{id} [delete]
func (h *TestHandler) DeleteTest(c *gin.Context) {
	role := c.Request.Context().Value(enum.ContextKeyRole)
	if role != enum.Moderator && role != enum.Admin {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	id, err := strconv.Atoi(mux.Vars(c.Request)["id"])
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.repo.DeleteTest(c, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.log.Error(err.Error())
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

// UpdateTest
// @Summary Обновление информации о тесте
// @Tags Test
// @Security Bearer
// @Param id path int true "Test ID"
// @Param input body dto.TestDTO true "Новая нформация о тесте"
// @Success 200
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /api/test/{id} [patch]
func (h *TestHandler) UpdateTest(c *gin.Context) {
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

	var testDTO dto.TestDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&testDTO); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.repo.UpdateTest(c, id, model.Test(testDTO))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.log.Error(err.Error())
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

// GetTest
// @Summary Получение теста по айди
// @Tags Test
// @Security Bearer
// @Param id path int true "Test ID"
// @Success 200 {object} dto.TestDTO
// @Failure 404
// @Failure 500
// @Router /api/test/{id} [get]
func (h *TestHandler) GetTest(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	test, err := h.repo.FindTestByID(c, id)
	if err != nil {
		h.log.Error(err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.TestDTO(test))
}

// GetTests
// @Summary Получение всех тестов
// @Tags Test
// @Security Bearer
// @Success 200 {object} []dto.TestDTO
// @Failure 404
// @Failure 500
// @Router /api/tests [get]
func (h *TestHandler) GetTests(c *gin.Context) {
	tests, err := h.repo.GetTests(c)
	if err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.TestsToTestsDTO(tests))
}

// GetTestsByTaskID
// @Summary Получение всех тестов для конкретной задачи
// @Tags Test
// @Security Bearer
// @Success 200  {object} []dto.TestDTO
// @Param task_id path int true "Task ID"
// @Failure 404
// @Failure 500
// @Router /api/tests/{task_id} [get]
func (h *TestHandler) GetTestsByTaskID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("task_id"))
	h.log.Info("GetTestsByTaskID recieved")

	if err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	h.log.Info("GetTestsByTaskID start")
	tests, err := h.repo.FindTestsByTaskID(c, taskID)
	if err != nil {
		h.log.Info("GetTestsByTaskID error")
		h.log.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	h.log.Info("GetTestsByTaskID end")

	c.JSON(http.StatusOK, dto.TestsToTestsDTO(tests))
}
