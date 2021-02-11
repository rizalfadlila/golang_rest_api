package repositories

import (
	"context"

	"github.com/rest_api/entities/models"
)

// CustomerRepository :nodoc:
type CustomerRepository interface {
	Synchronize(ctx context.Context) ([]models.Customer, error)
}
