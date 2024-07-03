package handler

import (
	"edusync/internal/model"
	"edusync/pkg/validator"

	"github.com/gin-gonic/gin"
)

// Create Lid
// @Description Create Lid
// @Summary Create Lid
// @Tags Lid
// @Accept json
// @Produce json
// @Param create body model.LidCreateRequest true "Create Lid"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/lids [post]
// @Security ApiKeyAuth
func (h *Handler) createLid(c *gin.Context) {
	var (
		err   error
		input model.LidCreateRequest
	)
	if err = c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	id, err := h.services.LidService.LidWriter.Create(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, Created, gin.H{
		"id": id,
	})
}

// Get List Lid
// @Description Get List Lid
// @Summary Get List Lid
// @Tags Lid
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/lids [get]
// @Security ApiKeyAuth
func (h *Handler) getListLid(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	lids, err := h.services.LidService.LidReader.GetList(&pagination)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, gin.H{
		"list":       lids,
		"pagination": pagination,
	})
}

// Get Lid By Id
// @Description Get Lid By Id
// @Summary Get Lid By Id
// @Tags Lid
// @Accept json
// @Produce json
// @Param id path int64 true "Lid Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/lids/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getLidById(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	lid, err := h.services.LidService.LidReader.GetById(id)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, lid)
}

// Update Lid
// @Description Update Lid
// @Summary Update Lid
// @Tags Lid
// @Accept json
// @Produce json
// @Param id path int64 true "Lid Id"
// @Param update body model.LidUpdateRequest true "Update Lid"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/lids/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateLid(c *gin.Context) {
	var input model.LidUpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}
	input.Id = id

	err = h.services.LidService.LidWriter.Update(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil)
}

// Delete Lid
// @Description Delete Lid
// @Summary Delete Lid
// @Tags Lid
// @Accept json
// @Produce json
// @Param id path int64 true "Lid Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/lids/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteLid(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	err = h.services.LidService.LidWriter.Delete(id)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil)
}
