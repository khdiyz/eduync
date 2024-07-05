package handler

import (
	"edusync/internal/model"
	"errors"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func successResponse(c *gin.Context, status Status, data interface{}, pagination *model.Pagination) {
	c.JSON(status.Code, model.BaseResponse{
		Success:     true,
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
		Pagination:  pagination,
	})
}

func errorResponse(c *gin.Context, status Status, err error) {
	c.JSON(status.Code, model.BaseResponse{
		Success:      false,
		Status:       status.Status,
		Description:  status.Description,
		ErrorMessage: err.Error(),
	})
}

func abortResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(Aborted.Code, model.BaseResponse{
		Success:      false,
		Status:       Aborted.Status,
		Description:  Aborted.Description,
		ErrorMessage: message,
	})
}

// CONVERT SERVICE ERROR TO HANDLER ERROR

func fromError(c *gin.Context, serviceError error) {
	st, _ := status.FromError(serviceError)
	err := st.Message()

	switch st.Code() {
	case codes.Internal:
		errorResponse(c, Internal, errors.New(err))
	case codes.NotFound:
		errorResponse(c, NotFound, errors.New(err))
	case codes.InvalidArgument:
		errorResponse(c, BadRequest, errors.New(err))
	case codes.Unavailable:
		errorResponse(c, Unavailable, errors.New(err))
	case codes.AlreadyExists:
		errorResponse(c, AlreadyExists, errors.New(err))
	case codes.Unauthenticated:
		errorResponse(c, Unauthorized, errors.New(err))
	}
}
