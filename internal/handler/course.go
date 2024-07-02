package handler

import (
	"edusync/internal/model"
	"edusync/pkg/response"
	"edusync/pkg/validator"

	"github.com/gin-gonic/gin"
)

// Create Course
// @Description Create Course
// @Summary Create Course
// @Tags Course
// @Accept json
// @Produce json
// @Param create body model.CourseCreateRequest true "Create Course"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/courses [post]
// @Security ApiKeyAuth
func (h *Handler) createCourse(c *gin.Context) {
	var (
		err   error
		input model.CourseCreateRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		response.Error(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.Error(c, response.BadRequest, err)
		return
	}

	id, err := h.services.CourseService.CourseWriter.Create(input)
	if err != nil {
		response.ServiceErrorConvert(c, err)
		return
	}

	response.Success(c, response.Created, gin.H{
		"id": id,
	})
}

// Get List Course
// @Description Get List Course
// @Summary Get List Course
// @Tags Course
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/courses [get]
// @Security ApiKeyAuth
func (h *Handler) getListCourse(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		response.Error(c, response.BadRequest, err)
		return
	}

	roles, err := h.services.CourseService.CourseReader.GetList(&pagination)
	if err != nil {
		response.ServiceErrorConvert(c, err)
		return
	}

	response.Success(c, response.OK, gin.H{
		"list":       roles,
		"pagination": pagination,
	})
}

// Get Course By Id
// @Description Get Course By Id
// @Summary Get Course By Id
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int64 true "Course Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/courses/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getCourseById(c *gin.Context) {
	id, err := getNullInt64Param(c, "id")
	if err != nil {
		response.Error(c, response.BadRequest, err)
		return
	}

	course, err := h.services.CourseService.CourseReader.GetById(id)
	if err != nil {
		response.ServiceErrorConvert(c, err)
		return
	}

	response.Success(c, response.OK, gin.H{
		"course": course,
	})
}

// Update Course
// @Description Update Course
// @Summary Update Course
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int64 true "Course Id"
// @Param update body model.CourseUpdateRequest true "Update Course"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/courses/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateCourse(c *gin.Context) {
	var input model.CourseUpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.Error(c, response.BadRequest, err)
		return
	}

	id, err := getNullInt64Param(c, "id")
	if err != nil {
		response.Error(c, response.BadRequest, err)
		return
	}
	input.Id = id

	err = h.services.CourseService.CourseWriter.Update(input)
	if err != nil {
		response.ServiceErrorConvert(c, err)
		return
	}

	response.Success(c, response.OK, nil)
}

// Delete Course
// @Description Delete Course
// @Summary Delete Course
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int64 true "Course Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/courses/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteCourse(c *gin.Context) {
	id, err := getNullInt64Param(c, "id")
	if err != nil {
		response.Error(c, response.BadRequest, err)
		return
	}

	err = h.services.CourseService.CourseWriter.Delete(id)
	if err != nil {
		response.ServiceErrorConvert(c, err)
		return
	}

	response.Success(c, response.OK, nil)
}
