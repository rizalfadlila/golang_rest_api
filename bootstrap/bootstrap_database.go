package bootstrap

import (
	mongo "github.com/rest_api/db/mongo"
	"github.com/rest_api/pkg/logger"
)

func initDatabases() {
	initMongoDB()
}

func initMongoDB() {
	mongoClient, err := mongo.GetClient()
	if err != nil {
		logger.Panic(
			logger.SetMessageFormat("Error at Bootstrap > initMongoDB: %v", err),
		)
	}
	mongoDatabase = mongo.GetDB(mongoClient, cfg.Databases.MongoDB.Database)
}
