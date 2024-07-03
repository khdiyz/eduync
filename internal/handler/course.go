package handler

import (
	"edusync/internal/model"
	"edusync/pkg/validator"

	"github.com/gin-gonic/gin"
)

var (
	idQuery = "id"
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
		errorResponse(c, BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	id, err := h.services.CourseService.CourseWriter.Create(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, Created, gin.H{
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
		errorResponse(c, BadRequest, err)
		return
	}

	roles, err := h.services.CourseService.CourseReader.GetList(&pagination)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, gin.H{
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
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	course, err := h.services.CourseService.CourseReader.GetById(id)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, gin.H{
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

	err = h.services.CourseService.CourseWriter.Update(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil)
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
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	err = h.services.CourseService.CourseWriter.Delete(id)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil)
}

// Create Course Exam Type
// @Description Create Course Exam Type
// @Summary Create Course Exam Type
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int64 true "Course Id"
// @Param create body model.ExamTypeCreateRequest true "Create Course"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/courses/{id}/exam-types [post]
// @Security ApiKeyAuth
func (h *Handler) createCourseExamType(c *gin.Context) {
	var (
		err   error
		input model.ExamTypeCreateRequest
	)

	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	if err = c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}
	input.CourseId = id

	if err := validator.ValidatePayloads(input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	id, err = h.services.CourseService.ExamTypeWriter.Create(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, Created, gin.H{
		"id": id,
	})
}

// Get List Course
// @Description Get List Course
// @Summary Get List Course
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int64 true "Course Id"
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/courses/{id}/exam-types [get]
// @Security ApiKeyAuth
func (h *Handler) getListCourseExamType(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	pagination, err := listPagination(c)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	examTypes, err := h.services.CourseService.ExamTypeReader.GetList(id, &pagination)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, gin.H{
		"list":       examTypes,
		"pagination": pagination,
	})
}

// Get List Course Exam Type
// @Description Get List Course Exam Type
// @Summary Get List Course Exam Type
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int64 true "Course Id"
// @Param examTypeId path int64 true "Exam Type Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/courses/{id}/exam-types/{examTypeId} [get]
// @Security ApiKeyAuth
func (h *Handler) getCourseExamType(c *gin.Context) {
	courseId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	examTypeId, err := getNullInt64Param(c, "examTypeId")
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	examType, err := h.services.CourseService.ExamTypeReader.GetById(courseId, examTypeId)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, examType)
}

// Update Course Exam Type
// @Description Update Course Exam Type
// @Summary Update Course Exam Type
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int64 true "Course Id"
// @Param examTypeId path int64 true "Exam Type Id"
// @Param update body model.ExamTypeUpdateRequest true "Update Course"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/courses/{id}/exam-types/{examTypeId} [put]
// @Security ApiKeyAuth
func (h *Handler) updateCourseExamType(c *gin.Context) {
	var (
		err   error
		input model.ExamTypeUpdateRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	input.CourseId, err = getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	input.Id, err = getNullInt64Param(c, "examTypeId")
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	err = h.services.CourseService.ExamTypeWriter.Update(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil)
}

// Delete Course Exam Type
// @Description Delete Course Exam Type
// @Summary Delete Course Exam Type
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int64 true "Course Id"
// @Param examTypeId path int64 true "Exam Type Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/courses/{id}/exam-types/{examTypeId} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteCourseExamType(c *gin.Context) {
	courseId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	examTypeId, err := getNullInt64Param(c, "examTypeId")
	if err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	err = h.services.CourseService.ExamTypeWriter.Delete(courseId, examTypeId)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, nil)
}
