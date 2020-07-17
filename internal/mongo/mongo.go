package mongo

import (
	"context"
	"github.com/lyeka/gotd/internal/UUID"
	"github.com/lyeka/gotd/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	CollUser     = "user"      // 用户集合
	CollUserTask = "user_task" // 用户任务集合
)

type DB struct {
	IDGenerator *UUID.IDGenerator // id生成器
	Client      *mongo.Client
	DB          *mongo.Database
}

func OpenDB(ctx context.Context, cfg *config.Config) (*DB, error) {
	dbName, dsn := cfg.DbDSN()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	idg, err := UUID.NewNode()
	if err != nil {
		return nil, err
	}

	return &DB{Client: client, DB: client.Database(dbName), IDGenerator: idg}, nil
}

func (db *DB) CollUser() *mongo.Collection {
	return db.DB.Collection(CollUser)
}

func (db *DB) CollTask() *mongo.Collection {
	return db.DB.Collection(CollUserTask)
}

func (db *DB) GenerateID() int64 {
	return db.IDGenerator.Generate().Int64()
}
