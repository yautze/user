package repository

import (
	"context"
	"user/app/domin/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepository -
type UserRepository interface {
	Create(ctx context.Context, in *model.User) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*model.User, error)
	GetByKey(ctx context.Context, key string) (*model.User, error)
	GetByFBID(ctx context.Context, fbID string) (*model.User, error)
	Find(ctx context.Context, filter interface{}) ([]*model.User, error)
	Update(ctx context.Context, filter, update interface{}) error
	DeleteByID(ctx context.Context, id primitive.ObjectID) error
}
