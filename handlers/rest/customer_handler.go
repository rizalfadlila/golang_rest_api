package rest

import (
	constantsREST "github.com/rest_api/constants/rest"
	"github.com/rest_api/usecases/services"

	"github.com/gin-gonic/gin"
)

// CustomerHandler :nodoc:
type CustomerHandler struct {
	baseHandler
	cutomerService services.CustomerService
}

// NewCustomerHandler :nodoc
func NewCustomerHandler(customerSvc services.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		cutomerService: customerSvc,
	}
}

// Synchronize godoc
// @Summary Synchronize Data
// @Description Synchronize data customer
// @Tags Transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.Customer string "Ok"
// @Failure 400 {string} responses.Response "Bad Request"
// @Failure 401 {string} responses.Response "Unauthorized"
// @Failure 500 {string} responses.Response "Internal Server Error"
// @Router /v1/customer/synchronize [get]
func (h *CustomerHandler) Synchronize(ginCtx *gin.Context) {

	customers, err := h.cutomerService.Synchronize(ginCtx.Request.Context())

	if err != nil {
		h.errorResponse(ginCtx, err, constantsREST.ResponseCodeInternalServerError)
		return
	}

	h.successResponse(ginCtx, customers, constantsREST.SuccessGetData)
}
