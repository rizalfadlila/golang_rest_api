package services

import (
	"context"
	"sync"

	"github.com/rest_api/entities/models"
	"github.com/rest_api/repositories"
)

var (
	customerSvc     CustomerService
	customerSvcOnce sync.Once
)

type cutomerService struct {
	cutomerRepository repositories.CustomerRepository
	defaultErrMsg     string
}

// NewCostumerService :nodoc:
func NewCostumerService(cutomerRepo repositories.CustomerRepository) CustomerService {
	customerSvcOnce.Do(func() {
		customerSvc = &cutomerService{
			cutomerRepository: cutomerRepo,
			defaultErrMsg:     "Error at CustomerService > ",
		}
	})
	return customerSvc
}

func (s *cutomerService) Synchronize(ctx context.Context) ([]models.Customer, error) {
	return s.cutomerRepository.Synchronize(ctx)
}
