package handler

import (
	"edusync/internal/model"
	"edusync/pkg/response"
	"edusync/pkg/validator"

	"github.com/gin-gonic/gin"
)

// Create User
// @Description Create User
// @Summary Create User
// @Tags User
// @Accept json
// @Produce json
// @Param create body model.UserCreateRequest true "Create User"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/users [post]
// @Security ApiKeyAuth
func (h *Handler) createUser(c *gin.Context) {
	var (
		err   error
		input model.UserCreateRequest
	)
	if err = c.ShouldBindJSON(&input); err != nil {
		response.Error(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.Error(c, response.BadRequest, err)
		return
	}

	id, err := h.services.UserWriter.Create(input)
	if err != nil {
		response.ServiceErrorConvert(c, err)
		return
	}

	response.Success(c, response.Created, gin.H{
		"id": id,
	})
}
