package services

import (
	"context"

	"github.com/rest_api/entities/models"
)

// CustomerService :nodoc:
type CustomerService interface {
	Synchronize(ctx context.Context) ([]models.Customer, error)
}
