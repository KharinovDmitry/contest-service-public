package handlers

import (
	"contest/internal/domain/enum"
	"contest/internal/domain/repository"
	"contest/internal/domain/service/logger"
	"contest/internal/server/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LaunchHandler struct {
	launchRepository repository.LaunchRepository
	log              logger.Logger
}

func NewLaunchHandler(launchRepository repository.LaunchRepository, log logger.Logger) *LaunchHandler {
	return &LaunchHandler{
		launchRepository: launchRepository,
		log:              log,
	}
}

// GetSuccessLaunchesByUser
// @Summary Получение результатов успешных решений задач конкретным пользователем
// @Tags Launch
// @Security Bearer
// @Param user_id path int true "User ID"
// @Success 200 {object} []dto.LaunchDTO
// @Failure 400
// @Failure 403
// @Failure 500
// @Router /api/launches/success/{user_id} [get]
func (h *LaunchHandler) GetSuccessLaunchesByUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		h.log.Warn(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	h.log.Info(strconv.Itoa(len(c.Keys)))
	user_id_ctx, exist := c.Get(enum.ContextKeyID)
	if !exist {
		h.log.Error("Ctx doesn't contains ID")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user_id_token, ok := user_id_ctx.(int)
	if !ok {
		h.log.Warn("Convert to int failed")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if user_id_token != userId {
		h.log.Warn("Incorrect id: excepted: %d received: %d", userId, user_id_token)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	launches, err := h.launchRepository.GetSuccessLaunchesByUser(c, userId)
	if err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.LaunchesToLaunchesDTO(launches))
}

// GetLaunchesByUser
// @Summary Получение результатов решения задач конкретным пользователем
// @Tags Launch
// @Security Bearer
// @Param user_id path int true "User ID"
// @Success 200 {object} []dto.LaunchDTO
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /api/launches/{user_id} [get]
func (h *LaunchHandler) GetLaunchesByUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user_id_ctx := c.Request.Context().Value(enum.ContextKeyID)
	user_id_token, ok := user_id_ctx.(int)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if user_id_token != userId {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	launches, err := h.launchRepository.GetLaunchesByUser(c, userId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, dto.LaunchesToLaunchesDTO(launches))
}

// GetLaunchesByUserAndContest
// @Summary Получение результатов решения конкретной задачи конкретным пользователем
// @Tags Launch
// @Security Bearer
// @Param user_id path int true "User ID"
// @Param contest_id path int true "Contest ID"
// @Success 200 {object} []dto.LaunchDTO
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /api/launches/{user_id}/{contest_id} [get]
func (h *LaunchHandler) GetLaunchesByUserAndContest(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user_id_ctx := c.Request.Context().Value(enum.ContextKeyID)
	user_id_token, ok := user_id_ctx.(int)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if user_id_token != userId {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	contestId, err := strconv.Atoi(c.Param("contest_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	launches, err := h.launchRepository.GetLaunchesByUserAndContest(c, userId, contestId)
	if err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.LaunchesToLaunchesDTO(launches))
}
