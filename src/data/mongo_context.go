package data

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	once          sync.Once
	mongoInstance IMongoContext
)

type MoncoContext struct {
	client *mongo.Client
	dbName string
}

func GetInstance() IMongoContext {
	once.Do(func() {
		mongoInstance = &MoncoContext{}
	})
	return mongoInstance
}

func (m *MoncoContext) Initialize(ctx context.Context, credential options.Credential, dbURI, dbName string, replicaset *string) error {
	clientOptions := options.Client()
	clientOptions.ReplicaSet = replicaset
	client, err := mongo.Connect(ctx, clientOptions.ApplyURI(dbURI).SetAuth(credential))
	if err != nil {
		return err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	m.dbName = dbName
	m.client = client
	return nil
}

func (m *MoncoContext) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	return m.client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			return err
		}
		err = fn(sessionContext)
		if err != nil {
			return sessionContext.AbortTransaction(sessionContext)
		}
		return sessionContext.CommitTransaction(sessionContext)
	})
}

// Insert stores documents in the collection
func (m *MoncoContext) Insert(ctx context.Context, collName string, doc interface{}) (string, error) {
	insertedObject, err := m.client.Database(m.dbName).Collection(collName).InsertOne(ctx, doc)
	if insertedObject == nil {
		return "", err
	}

	id := insertedObject.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

// BulkInsert stores multiple documents in the collection
func (m *MoncoContext) BulkInsert(ctx context.Context, collName string, docs []interface{}) ([]interface{}, error) {
	insertedObject, err := m.client.Database(m.dbName).Collection(collName).InsertMany(ctx, docs)
	if insertedObject == nil {
		return nil, err
	}
	return insertedObject.InsertedIDs, err
}

// Find finds all documents in the collection
func (m *MoncoContext) Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}, opts *options.FindOptions) error {
	cur, err := m.client.Database(m.dbName).Collection(collName).Find(ctx, query, opts)
	if err != nil {
		return err
	}

	if err = cur.All(ctx, doc); err != nil {
		return err
	}

	return nil
}

// FindOne finds one document in mongo
func (m *MoncoContext) FindOne(ctx context.Context, collName string, query map[string]interface{}, doc interface{}, opts *options.FindOneOptions) error {
	return m.client.Database(m.dbName).Collection(collName).FindOne(ctx, query, opts).Decode(doc)
}

// UpdateOne updates one or more documents in the collection
func (m *MoncoContext) UpdateOne(ctx context.Context, collName string, selector map[string]interface{}, update interface{}) (*mongo.UpdateResult, error) {
	updateResult, err := m.client.Database(m.dbName).Collection(collName).UpdateOne(ctx, selector, update)
	return updateResult, err
}

// UpdateMany updates one or more documents in the collection
func (m *MoncoContext) UpdateMany(ctx context.Context, collName string, selector map[string]interface{}, update interface{}) (*mongo.UpdateResult, error) {
	updateResult, err := m.client.Database(m.dbName).Collection(collName).UpdateMany(ctx, selector, update)
	return updateResult, err
}

// Remove one documents in the collection
func (m *MoncoContext) Remove(ctx context.Context, collName string, selector map[string]interface{}) error {
	_, err := m.client.Database(m.dbName).Collection(collName).DeleteOne(ctx, selector)
	return err
}

// Remove many documents in the collection
func (m *MoncoContext) RemoveMany(ctx context.Context, collName string, selector map[string]interface{}) error {
	_, err := m.client.Database(m.dbName).Collection(collName).DeleteMany(ctx, selector)
	return err
}

// Count returns the number of documents of the query
func (m *MoncoContext) Count(ctx context.Context, collName string, query map[string]interface{}) (int64, error) {
	return m.client.Database(m.dbName).Collection(collName).CountDocuments(ctx, query)
}

func (m *MoncoContext) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_ = m.client.Disconnect(ctx)
}
