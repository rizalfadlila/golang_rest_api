package rest

import (
	"fmt"
	"net/http"

	constantsrest "github.com/rest_api/constants/rest"
	"github.com/rest_api/handlers/rest/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type baseHandler struct {
	validate *validator.Validate
}

// SuccessResponse :nodoc:
func (handler *baseHandler) successResponse(ginCtx *gin.Context, data interface{}, message string) {
	ginCtx.JSON(http.StatusOK, responses.Response{
		Data:    data,
		Message: message,
		Status:  constantsrest.StatusSuccess,
	})
}

// response :nodoc:
func (handler *baseHandler) response(ginCtx *gin.Context, data interface{}) {
	ginCtx.JSON(http.StatusOK, data)
}

// ErrorResponse :nodoc:
func (handler *baseHandler) errorResponse(ginCtx *gin.Context, err error, errorCode string) {
	errorConstant := responses.GetErrorConstant(errorCode)

	ginCtx.JSON(errorConstant.HTTPCode, responses.Response{
		Errors:  parseError(err),
		Message: errorConstant.Message,
		Status:  constantsrest.StatusFailed,
	})
}

// NotFoundResponse :nodoc:
func (handler *baseHandler) notFoundResponse(ginCtx *gin.Context) {
	errorConstant := responses.GetErrorConstant(constantsrest.ResponseCodeDataNotFound)

	ginCtx.JSON(errorConstant.HTTPCode, responses.Response{
		Message: errorConstant.Message,
		Status:  constantsrest.StatusFailed,
	})
}

func (handler *baseHandler) setDefaultParamPagination(ginCtx *gin.Context, params *map[string]interface{}) {
	*params = map[string]interface{}{
		constantsrest.Page:  ginCtx.Request.Header.Get(constantsrest.Page),
		constantsrest.Limit: ginCtx.Request.Header.Get(constantsrest.Limit),
	}
}

// paginationResponse :nodoc:
func (handler *baseHandler) paginationResponse(ginCtx *gin.Context, data responses.Pagination, message string) {
	ginCtx.JSON(http.StatusOK, responses.Response{
		Data:    data,
		Message: message,
		Status:  constantsrest.StatusSuccess,
	})
}

// parseError :nodoc:
func parseError(err error) []string {
	if err == nil {
		return nil
	}

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return []string{err.Error()}
	}

	var errors []string
	for _, e := range ve {
		errors = append(errors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", e.Field(), e.Tag()))
	}

	return errors
}
