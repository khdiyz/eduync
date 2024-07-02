package response

import (
	"edusync/internal/model"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SERVICE ERROR

var errorMapping = map[string]struct {
	code codes.Code
	msg  string
}{
	"no rows in result set":                          {codes.NotFound, "data is empty"},
	"duplicate key value violates unique constraint": {codes.AlreadyExists, "variable value is already exists"},
	"violates foreign key constraint":                {codes.InvalidArgument, "foreign key violation"},
	"no rows affected":                               {codes.Internal, "variable value is not exists"},
}

func ServiceError(err error, code codes.Code) error {
	if err == nil {
		return nil
	}

	errMsg := err.Error()

	for substr, mapping := range errorMapping {
		if strings.Contains(errMsg, substr) {
			return status.Error(mapping.code, mapping.msg)
		}
	}

	if code != codes.OK {
		return status.Error(code, errMsg)
	}

	return status.Error(codes.Unknown, errMsg)
}

// HANDLER RESPONSE

func Success(c *gin.Context, status Status, data interface{}) {
	c.JSON(status.Code, model.BaseResponse{
		Success:     true,
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

func Error(c *gin.Context, status Status, err error) {
	c.JSON(status.Code, model.BaseResponse{
		Success:      false,
		Status:       status.Status,
		Description:  status.Description,
		ErrorMessage: err.Error(),
	})
}

func Abort(c *gin.Context, message string) {
	c.AbortWithStatusJSON(Aborted.Code, model.BaseResponse{
		Success:      false,
		Status:       Aborted.Status,
		Description:  Aborted.Description,
		ErrorMessage: message,
	})
}

// CONVERT SERVICE ERROR TO HANDLER ERROR

func ServiceErrorConvert(c *gin.Context, serviceError error) {
	st, _ := status.FromError(serviceError)
	err := st.Message()

	switch st.Code() {
	case codes.Internal:
		Error(c, Internal, nil)
	case codes.NotFound:
		Error(c, NotFound, errors.New(err))
	case codes.InvalidArgument:
		Error(c, BadRequest, errors.New(err))
	case codes.Unavailable:
		Error(c, Unavailable, errors.New(err))
	case codes.AlreadyExists:
		Error(c, AlreadyExists, errors.New(err))
	case codes.Unauthenticated:
		Error(c, Unauthorized, errors.New(err))
	}
}
