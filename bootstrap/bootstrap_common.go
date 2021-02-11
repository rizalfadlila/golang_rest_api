package bootstrap

import (
	"github.com/rest_api/config"
	"github.com/rest_api/repositories"
	"github.com/rest_api/usecases/services"

	"go.mongodb.org/mongo-driver/mongo"
)

var cfg = *config.NewConfig()

// databases
var mongoDatabase *mongo.Database

// repositories
var customerRepository repositories.CustomerRepository

// services
var customerService services.CustomerService
