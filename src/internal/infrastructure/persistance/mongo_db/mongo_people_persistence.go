package mongo_db

import (
	"context"
	"errors"
	"time"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoPeoplePersistence struct {
	database       *mongo.Database
	collectionName string
}

func NewMongoPeoplePersistence(db *mongo.Database) MongoPeoplePersistence {
	return MongoPeoplePersistence{
		database:       db,
		collectionName: "people",
	}
}

func (p MongoPeoplePersistence) FindOne(id string) (models.People, error) {
	collection := p.database.Collection(p.collectionName)

	var people models.People

	// Por esto:
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return people, domain_err.ErrInvalidObjectId
	}

	filter := bson.M{"_id": objectID}

	err = collection.FindOne(context.TODO(), filter).Decode(&people)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return people, domain_err.ErrNotFound
		}
		return people, err
	}

	return people, nil
}

func (p MongoPeoplePersistence) FindAll() ([]models.People, error) {
	collection := p.database.Collection(p.collectionName)

	var people []models.People

	filter := bson.M{}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return people, err
	}

	err = cursor.All(context.TODO(), &people)
	if err != nil {
		return people, err
	}

	return people, nil
}

func (p MongoPeoplePersistence) Create(people *models.People) error {

	collection := p.database.Collection(p.collectionName)

	people.ID = bson.NewObjectID()
	people.CreatedAt = time.Now()
	people.UpdatedAt = time.Now()

	_, err := collection.InsertOne(context.TODO(), people)
	if err != nil {

		var writeErr mongo.WriteException
		if errors.As(err, &writeErr) {
			for _, e := range writeErr.WriteErrors {
				if e.Code == 11000 {
					return domain_err.ErrDuplicatedDoc
				}
			}
		}

		return err
	}

	return nil
}

func (p MongoPeoplePersistence) Update(people *models.People) error {
	collection := p.database.Collection(p.collectionName)

	people.UpdatedAt = time.Now()

	filter := bson.M{"_id": people.ID}
	update := bson.M{"$set": people}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
