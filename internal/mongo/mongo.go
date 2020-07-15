package mongo

import (
	"context"
	"github.com/lyeka/gotd/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	CollUser = "user"
)


type DB struct {
	DB *mongo.Database
}

func OpenDB(ctx context.Context, cfg *config.Config) (*DB, error){
	dbName, dsn := cfg.DbDSN()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &DB{DB: client.Database(dbName)}, nil
}

func (db *DB) CollUser() *mongo.Collection{
	return db.DB.Collection(CollUser)
}

