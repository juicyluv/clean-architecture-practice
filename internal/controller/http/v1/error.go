package v1

import (
	"clean-arch/internal/usecase/apperror"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func errorResponse(c *gin.Context, err error) {
	var (
		appError apperror.AppError
		code     int
	)

	if !errors.As(err, &appError) {
		appError.Message = err.Error()
	}

	switch appError.Type {
	case apperror.ErrorTypeNotFound:
		code = http.StatusNotFound
	case apperror.ErrorTypeInvalidRequest:
		code = http.StatusBadRequest
	default:
		code = http.StatusInternalServerError
	}

	c.AbortWithStatusJSON(code, appError)
}
