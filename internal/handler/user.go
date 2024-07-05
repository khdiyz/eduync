package handler

import (
	"edusync/internal/model"
	"edusync/pkg/validator"

	"github.com/gin-gonic/gin"
)

var (
	roleIdQuery = "role-id"
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
		errorResponse(c, BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	id, err := h.services.UserWriter.Create(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, Created, id, nil)
}

// Get List User
// @Description Get List User
// @Summary Get List User
// @Tags User
// @Accept json
// @Produce json
// @Param role-id query int64 false "filter by role id"
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/users [get]
// @Security ApiKeyAuth
func (h *Handler) getListUser(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	filter := make(map[string]interface{})

	roleId, err := getNullInt64Query(c, roleIdQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}
	if roleId != 0 {
		filter[roleIdQuery] = roleId
	}

	users, err := h.services.UserService.UserReader.GetList(&pagination, filter)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, users, &pagination)
}
