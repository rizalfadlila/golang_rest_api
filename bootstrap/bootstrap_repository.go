package bootstrap

import "github.com/rest_api/repositories"

func initRepository() {
	customerRepository = repositories.NewCustomerRepository(mongoDatabase)
}
