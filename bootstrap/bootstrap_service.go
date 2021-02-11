package bootstrap

import "github.com/rest_api/usecases/services"

func initService() {
	customerService = services.NewCostumerService(customerRepository)
}
