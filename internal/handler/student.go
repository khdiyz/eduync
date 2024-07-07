package handler

import (
	"edusync/internal/model"
	"edusync/pkg/validator"

	"github.com/gin-gonic/gin"
)

// Create Student
// @Description Create Student
// @Summary Create Student
// @Tags Student
// @Accept json
// @Produce json
// @Param create body model.StudentCreateRequest true "Create Student"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/students [post]
// @Security ApiKeyAuth
func (h *Handler) createStudent(c *gin.Context) {
	var (
		err   error
		input model.StudentCreateRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	id, err := h.services.StudentService.StudentWriter.Create(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, Created, id, nil)
}

// Get List Student
// @Description Get List Student
// @Summary Get List Student
// @Tags Student
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/students [get]
// @Security ApiKeyAuth
func (h *Handler) getListStudent(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	students, err := h.services.StudentService.StudentReader.GetList(&pagination)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, students, &pagination)
}

// Get Student By Id
// @Description Get Student By Id
// @Summary Get Student By Id
// @Tags Student
// @Accept json
// @Produce json
// @Param id path int64 true "Student Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/students/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getStudentById(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	student, err := h.services.StudentService.StudentReader.GetById(id)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, student, nil)
}

// Update Student
// @Description Update Student
// @Summary Update Student
// @Tags Student
// @Accept json
// @Produce json
// @Param id path int64 true "Student Id"
// @Param update body model.StudentUpdateRequest true "Update Student"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/students/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateStudent(c *gin.Context) {
	var input model.StudentUpdateRequest

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

	err = h.services.StudentService.StudentWriter.Update(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil, nil)
}

// Delete Student
// @Description Delete Student
// @Summary Delete Student
// @Tags Student
// @Accept json
// @Produce json
// @Param id path int64 true "Student Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/students/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteStudent(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	err = h.services.StudentService.StudentWriter.Delete(id)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil, nil)
}

// Get List Student Actions
// @Description Get List Student Actions
// @Summary Get List Student Actions
// @Tags Student
// @Accept json
// @Produce json
// @Param id path int64 true "Student Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/students/{id}/actions [get]
// @Security ApiKeyAuth
func (h *Handler) getListStudentActions(c *gin.Context) {
	studentId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	actions, err := h.services.StudentService.StudentReader.GetActions(studentId)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, actions, nil)
}
