package repositories

import (
	"context"

	"github.com/rest_api/entities/models"
	"github.com/rest_api/helpers"
	"github.com/rest_api/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type customerRepository struct {
	db             *mongo.Database
	collectionName string
	defaultErrMsg  string
}

// NewCustomerRepository :nodoc:
func NewCustomerRepository(db *mongo.Database) CustomerRepository {
	return &customerRepository{
		db:             db,
		collectionName: "customer",
		defaultErrMsg:  "Error at CustomerRepository > ",
	}
}

func (r *customerRepository) Synchronize(ctx context.Context) ([]models.Customer, error) {
	var err error = nil
	customers := make([]models.Customer, 0)

	helpers.Block{
		Try: func() {
			cursor, err := r.db.Collection(r.collectionName).Find(ctx, bson.D{})
			helpers.Throw(err)

			defer cursor.Close(ctx)

			for cursor.Next(ctx) {
				var customer models.Customer
				helpers.Throw(cursor.Decode(&customer))
				customers = append(customers, customer)
			}
		},
		Catch: func(e helpers.Exception) {
			logger.Error(logger.SetMessageFormat("%v Synchronize: %v", r.defaultErrMsg, e))
			err = e
		},
	}.Do()

	return customers, err
}
