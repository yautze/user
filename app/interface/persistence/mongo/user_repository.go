package mongo

import (
	"context"
	"user/app/domin/model"
	"user/config"

	mgo "github.com/yautze/tools/db/mongo"
	"github.com/yautze/tools/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// COLLECTIONUSER -
const COLLECTIONUSER = "DDDUser"

type userRepository struct {
	database string
	client   func() (*mongo.Client, error)
}

// NewUserRepository -
func NewUserRepository(m config.Mongo) *userRepository {
	// connect mongo
	connect(m)

	return &userRepository{
		database: m.Database,
		client:   mgo.M,
	}
}

// Create -
func (u *userRepository) Create(ctx context.Context, in *model.User) error {
	client, err := u.client()
	if err != nil {
		logger.WithError(err).Error("connect mongo failed")
		return err
	}

	// connection  database and collection
	coll := client.Database(u.database).Collection(COLLECTIONUSER)

	if _, err := coll.InsertOne(ctx, in); err != nil {
		logger.WithError(err).Error("insert user failed")
		return err
	}

	return nil
}

// GetByID -
func (u *userRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*model.User, error) {
	client, err := u.client()
	if err != nil {
		logger.WithError(err).Error("connect mongo failed")
		return nil, err
	}

	// connection  database and collection
	coll := client.Database(u.database).Collection(COLLECTIONUSER)

	filter := bson.M{
		"_id": id,
	}

	var res *model.User
	if err := coll.FindOne(ctx, filter).Decode(&res); err != nil {
		logger.WithError(err).Error("get user by id failed")
		return nil, err
	}

	return res, nil
}

// GetByKey -
func (u *userRepository) GetByKey(ctx context.Context, key string) (*model.User, error) {
	client, err := u.client()
	if err != nil {
		logger.WithError(err).Error("connect mongo failed")
		return nil, err
	}

	// connection  database and collection
	coll := client.Database(u.database).Collection(COLLECTIONUSER)

	filter := bson.M{
		"key": key,
	}

	var res *model.User
	if err := coll.FindOne(ctx, filter).Decode(&res); err != nil {
		logger.WithError(err).Error("get user by key failed")
		return nil, err
	}

	return res, nil
}

// GetByFBID -
func (u *userRepository) GetByFBID(ctx context.Context, fbID string) (*model.User, error) {
	client, err := u.client()
	if err != nil {
		logger.WithError(err).Error("connect mongo failed")
		return nil, err
	}

	// connection  database and collection
	coll := client.Database(u.database).Collection(COLLECTIONUSER)

	filter := bson.M{
		"fbID": fbID,
	}

	var res *model.User
	if err := coll.FindOne(ctx, filter).Decode(&res); err != nil {
		logger.WithError(err).Error("get user by key failed")
		return nil, err
	}

	return res, nil
}

// Find -
func (u *userRepository) Find(ctx context.Context, filter interface{}) ([]*model.User, error) {
	client, err := u.client()
	if err != nil {
		logger.WithError(err).Error("connect mongo failed")
		return nil, err
	}

	// connection  database and collection
	coll := client.Database(u.database).Collection(COLLECTIONUSER)
	cur, err := coll.Find(ctx, filter)

	if err != nil {
		logger.WithError(err).Error("find keyword failed")
		return nil, err
	}

	var res []*model.User
	if err := cur.All(ctx, &res); err != nil {
		logger.WithError(err).Error("cursor keyword failed")
		return nil, err
	}

	return res, nil
}

// Update -
func (u *userRepository) Update(ctx context.Context, filter, update interface{}) error {
	client, err := u.client()
	if err != nil {
		logger.WithError(err).Error("connect mongo failed")
		return err
	}

	// connection  database and collection
	coll := client.Database(u.database).Collection(COLLECTIONUSER)

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		logger.WithError(err).Error("update user failed")
		return err
	}

	return nil
}

// DeleteByID -
func (u *userRepository) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	client, err := u.client()
	if err != nil {
		logger.WithError(err).Error("connect mongo failed")
		return err
	}

	// connection  database and collection
	coll := client.Database(u.database).Collection(COLLECTIONUSER)

	filter := bson.M{
		"_id": id,
	}

	if _, err := coll.DeleteOne(ctx, filter); err != nil {
		logger.WithError(err).Error("delete user by id failed")
		return err
	}

	return nil
}
