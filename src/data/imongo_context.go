package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoContext interface {
	Insert(ctx context.Context, collName string, doc interface{}) (string, error)
	BulkInsert(ctx context.Context, collName string, docs []interface{}) ([]interface{}, error)
	Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}, opts *options.FindOptions) error
	FindOne(ctx context.Context, collName string, query map[string]interface{}, doc interface{}, opts *options.FindOneOptions) error
	Count(ctx context.Context, collName string, query map[string]interface{}) (int64, error)
	UpdateOne(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, collName string, selector map[string]interface{}, update interface{}) (*mongo.UpdateResult, error)
	Remove(ctx context.Context, collName string, query map[string]interface{}) error
	RemoveMany(ctx context.Context, collName string, selector map[string]interface{}) error
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
	Initialize(ctx context.Context, credential options.Credential, dbURI string, dbName string, replicaset *string) error
	Disconnect()
}
