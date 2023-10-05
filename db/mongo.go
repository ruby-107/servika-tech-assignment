// db/mongodb.go
package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *mongo.Client
	Database *mongo.Database
)

func InitMongoDB(connectionString, dbName string) error {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(connectionString)
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = c.Ping(ctx, nil)
	if err != nil {
		return err
	}

	client = c
	Database = client.Database(dbName)
	return nil
}
