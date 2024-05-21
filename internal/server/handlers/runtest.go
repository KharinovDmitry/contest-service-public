package handlers

import (
	"contest/internal/domain/enum"
	"contest/internal/domain/service/executor"
	"contest/internal/domain/service/logger"
	"contest/internal/domain/service/testRunner"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RunError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RunTestRequest struct {
	TaskID int `json:"task_id"`
	// @Model enum.Language
	Language enum.Language `json:"language"`
	Code     string        `json:"code"`
}

type RunTestResponse struct {
	//@Model enum.TestResultCode
	ResultCode  string `json:"result_code"`
	Description string `json:"description"`
	Points      int    `json:"points,string"`
}

type RunTestHandler struct {
	runTestService testRunner.TestRunner
	log            logger.Logger
}

func NewRunTestHandler(runTestService testRunner.TestRunner, log logger.Logger) *RunTestHandler {
	return &RunTestHandler{
		runTestService: runTestService,
		log:            log,
	}
}

// RunTest
// @Summary Отправка кода на проверку
// @Tags Run Test
// @Security Bearer
// @Accept json
// @Param input body handlers.RunTestRequest true "инормация о проверяемом коде"
// @Success 200 {object} handlers.RunTestResponse
// @Failure 403 {object} RunError
// @Failure 404 {object} RunError
// @Failure 500 {object} RunError
// @Router /api/run [POST]
func (h *RunTestHandler) RunTest(c *gin.Context) {
	var request RunTestRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, RunError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	user_id_ctx := c.Request.Context().Value(enum.ContextKeyID)
	user_id, ok := user_id_ctx.(int)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, RunError{
			Code:    http.StatusBadRequest,
			Message: "error getting user id",
		})
		return
	}

	result, err := h.runTestService.RunTest(c, request.TaskID, user_id, request.Language, request.Code)
	if err != nil {
		if errors.Is(err, executor.ErrUnknownLanguage) {
			c.AbortWithStatusJSON(http.StatusBadRequest, RunError{
				Code:    http.StatusBadRequest,
				Message: executor.ErrUnknownLanguage.Error(),
			})
			return

		}
		h.log.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, RunTestResponse{
		ResultCode:  string(result.ResultCode),
		Description: result.Description,
		Points:      result.Points,
	})
}
