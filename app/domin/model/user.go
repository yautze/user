package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User -
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	FBID     string             `bson:"fbID"`
	Password string             `bson:"-"`
	Key      string             `bson:"key"`
	Info     Info               `bson:"info"`
	CreateAt int64              `bson:"createAt"`
	UpdateAt int64              `bson:"updateAt"`
	LoginAt  int64              `bson:"loginAt"`
}

// Info - user info
type Info struct {
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
}
