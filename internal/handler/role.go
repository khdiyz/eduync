package handler

import (
	"edusync/pkg/response"

	"github.com/gin-gonic/gin"
)

// GetListRole
// @Description Get List Role
// @Summary Get List Role
// @Tags Role
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/roles [get]
// @Security ApiKeyAuth
func (h *Handler) getListRole(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		response.Error(c, response.BadRequest, err)
		return
	}

	roles, err := h.services.RoleReader.GetList(&pagination)
	if err != nil {
		response.ServiceErrorConvert(c, err)
		return
	}

	response.Success(c, response.OK, gin.H{
		"list":       roles,
		"pagination": pagination,
	})
}
