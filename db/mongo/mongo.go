package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	config "github.com/rest_api/config"
	log "github.com/sirupsen/logrus"
	migrate "github.com/xakep666/mongo-migrate"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// GetClient :nodoc:
func GetClient() (*mongo.Client, error) {
	cfg := *config.GetConfig()
	var ConnectionURI string

	if cfg.App.Env == "production" {
		ConnectionURI = cfg.Databases.MongoDB.URI
	} else {
		ConnectionURI = fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=%s",
			cfg.Databases.MongoDB.Authentication.Username,
			cfg.Databases.MongoDB.Authentication.Password,
			cfg.Databases.MongoDB.Host,
			cfg.Databases.MongoDB.Database,
			cfg.Databases.MongoDB.Authentication.Database,
		)
	}

	client, err := mongo.NewClient(options.Client().SetAuth(
		options.Credential{
			Username:   cfg.Databases.MongoDB.Authentication.Username,
			Password:   cfg.Databases.MongoDB.Authentication.Password,
			AuthSource: cfg.Databases.MongoDB.Authentication.Database,
		}).ApplyURI(ConnectionURI))

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Warn(fmt.Sprintf("Failed connected to MongoDB %s", ConnectionURI))
		return nil, errors.Wrap(err, fmt.Sprintf("test Failed to connect MongoDB %s", ConnectionURI))
	}

	log.Info(fmt.Sprintf("Successfully connected to database %s", ConnectionURI))

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, errors.Wrap(err, "Could not ping to mongodb server")
	}

	return client, nil
}

// GetDB :noodoc:
func GetDB(client *mongo.Client, dbName string) *mongo.Database {
	DB := client.Database(dbName)
	return DB
}

// Ping :nodoc:
func Ping(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return errors.Wrap(err, "Could not ping to mongodb server")
	}

	return nil
}

// Migrate :nodoc:
func Migrate() error {
	cfg := *config.GetConfig()

	client, err := GetClient()
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to connect MongoDB Client"))
	}

	defer client.Disconnect(context.Background())

	db := GetDB(client, cfg.Databases.MongoDB.Database)
	migrate.SetDatabase(db)
	if err := migrate.Up(migrate.AllAvailable); err != nil {
		return err
	}

	return nil
}
