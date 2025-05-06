package repository

import (
	"be-fiber/database"
	"be-fiber/model"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserByUsername(username string) (*model.User, error)
	ExistsByUsernameOrEmail(username, email string) (bool, error)
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *database.MongoDB
}

func NewUserRepository(db *database.MongoDB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.db.GetCollection("users")
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) ExistsByUsernameOrEmail(username, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}

	count, err := r.db.GetCollection("users").CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil 
}

func (r *userRepository) CreateUser(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	result, err := r.db.GetCollection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}

	return nil
}