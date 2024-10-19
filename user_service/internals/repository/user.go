package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"user_service/internals/db"
	"user_service/internals/models"
)

type User interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, ID int) (*models.User, error)
}

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(mongoClient *db.MongoClient, collectionName string) *UserRepository {
	return &UserRepository{collection: mongoClient.Database.Collection(collectionName)}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	user.ID = primitive.NewObjectID()

	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, ID int) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
