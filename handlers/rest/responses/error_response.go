package responses

import (
	"errors"
	"net/http"

	constantsrest "github.com/rest_api/constants/rest"
)

// ErrorConstant :nodoc:
type ErrorConstant struct {
	HTTPCode int
	Message  string
}

var errorConstantMapping = map[string]ErrorConstant{
	constantsrest.ResponseCodeBadRequest: {
		HTTPCode: http.StatusBadRequest,
		Message:  "Invalid request parameter",
	},
	constantsrest.ResponseCodeInternalServerError: {
		HTTPCode: http.StatusInternalServerError,
		Message:  "Something went wrong",
	},
	constantsrest.ResponseCodeInvalidRequest: {
		HTTPCode: http.StatusUnprocessableEntity,
		Message:  "Invalid Request",
	},
	constantsrest.ResponseCodeConflict: {
		HTTPCode: http.StatusConflict,
		Message:  "Request conflict",
	},
	constantsrest.ResponseCodeUnauthorized: {
		HTTPCode: http.StatusUnauthorized,
		Message:  "Unauthorized data",
	},
	constantsrest.ResponseCodeDataNotFound: {
		HTTPCode: http.StatusNotFound,
		Message:  "Resource not found",
	},
	constantsrest.ResponseCodeErrUnknown: {
		HTTPCode: http.StatusInternalServerError,
		Message:  "Error Code Unknown",
	},
}

// GetErrorConstant error message and http code
func GetErrorConstant(code string) ErrorConstant {
	return errorConstantMapping[code]
}

// ErrorInvalidID :nodoc:
func ErrorInvalidID() error {
	return errors.New("Invalid Id")
}

// ErrorIDRequired :nodoc:
func ErrorIDRequired() error {
	return errors.New("Id required")
}
