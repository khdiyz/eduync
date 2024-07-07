package handler

import (
	"edusync/internal/model"
	"edusync/pkg/validator"

	"github.com/gin-gonic/gin"
)

// Create Group
// @Description Create Group
// @Summary Create Group
// @Tags Group
// @Accept json
// @Produce json
// @Param create body model.GroupCreateRequest true "Create Group"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/groups [post]
// @Security ApiKeyAuth
func (h *Handler) createGroup(c *gin.Context) {
	var (
		err   error
		input model.GroupCreateRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	id, err := h.services.GroupService.GroupWriter.Create(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, Created, id, nil)
}

// Get List Group
// @Description Get List Group
// @Summary Get List Group
// @Tags Group
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/groups [get]
// @Security ApiKeyAuth
func (h *Handler) getListGroup(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	groups, err := h.services.GroupService.GroupReader.GetList(&pagination)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, groups, &pagination)
}

// Get Group By Id
// @Description Get Group By Id
// @Summary Get Group By Id
// @Tags Group
// @Accept json
// @Produce json
// @Param id path int64 true "Group Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/groups/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getGroupById(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	group, err := h.services.GroupService.GroupReader.GetById(id)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, group, nil)
}

// Update Group
// @Description Update Group
// @Summary Update Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path int64 true "Group Id"
// @Param update body model.GroupUpdateRequest true "Update Group"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/groups/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateGroup(c *gin.Context) {
	var input model.GroupUpdateRequest

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

	err = h.services.GroupService.GroupWriter.Update(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil, nil)
}

// Delete Group
// @Description Delete Group
// @Summary Delete Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path int64 true "Group Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/groups/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteGroup(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	err = h.services.GroupService.GroupWriter.Delete(id)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil, nil)
}

// Join Student to the Group
// @Description Join Student to the Group
// @Summary Join Student to the Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path int64 true "Group Id"
// @Param studentId path int64 true "Student Id"
// @Param joinDate body model.JoinStudentRequest true "Join Student"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/groups/{id}/join/{studentId} [post]
// @Security ApiKeyAuth
func (h *Handler) joinStudent(c *gin.Context) {
	var input model.JoinStudentRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	groupId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}
	input.GroupId = groupId

	studentId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}
	input.StudentId = studentId

	err = h.services.GroupService.GroupWriter.JoinStudent(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil, nil)
}

// Left Student from the Group
// @Description Left Student from the Group
// @Summary Left Student from the Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path int64 true "Group Id"
// @Param studentId path int64 true "Student Id"
// @Param leftDate body model.LeftStudentRequest true "Left Student"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/groups/{id}/left/{studentId} [put]
// @Security ApiKeyAuth
func (h *Handler) leftStudent(c *gin.Context) {
	var input model.LeftStudentRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	groupId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}
	input.GroupId = groupId

	studentId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}
	input.StudentId = studentId

	err = h.services.GroupService.GroupWriter.LeftStudent(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil, nil)
}

// Freeze Student from Group
// @Description Freeze Student from Group
// @Summary Freeze Student from Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path int64 true "Group Id"
// @Param studentId path int64 true "Student Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/groups/{id}/freeze/{studentId} [put]
// @Security ApiKeyAuth
func (h *Handler) freezeStudent(c *gin.Context) {
	var input model.FreezeStudentRequest

	groupId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}
	input.GroupId = groupId

	studentId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}
	input.StudentId = studentId

	err = h.services.GroupService.GroupWriter.FreezeStudent(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil, nil)
}

// Unfreeze Student Group
// @Description Unfreeze Student Group
// @Summary Unfreeze Student Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path int64 true "Group Id"
// @Param studentId path int64 true "Student Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/groups/{id}/unfreeze/{studentId} [put]
// @Security ApiKeyAuth
func (h *Handler) unfreezeStudent(c *gin.Context) {
	var input model.UnfreezeStudentRequest

	groupId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}
	input.GroupId = groupId

	studentId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}
	input.StudentId = studentId

	err = h.services.GroupService.GroupWriter.UnfreezeStudent(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil, nil)
}
