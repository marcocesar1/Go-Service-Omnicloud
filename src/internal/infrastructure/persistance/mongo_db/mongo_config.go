package mongo_db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConfig struct {
	DbUrl  string
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoConfig(dbUrl string) *MongoConfig {
	return &MongoConfig{
		DbUrl: dbUrl,
	}
}

func (m *MongoConfig) Connect() error {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(m.DbUrl).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return err
	}

	if err := client.Database("admin").
		RunCommand(context.TODO(), bson.D{{"ping", 1}}).
		Err(); err != nil {
		return err
	}

	m.client = client

	fmt.Println("Successfully connected to MongoDB!")
	return nil
}

func (m *MongoConfig) Disconnect() error {
	if m.client == nil {
		return nil
	}
	return m.client.Disconnect(context.TODO())
}

func (m *MongoConfig) CreateCollections() {
	m.db = m.client.Database("admin")

	err := m.db.CreateCollection(context.TODO(), "people")
	if err != nil {
		log.Fatalf("Failed to create people collection: %v", err)
	}
}

func (m *MongoConfig) GetDatabase() *mongo.Database {
	return m.db
}
