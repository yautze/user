package mongo

import (
	"context"
	"time"
	"user/config"

	mgo "github.com/yautze/tools/db/mongo"
	"github.com/yautze/tools/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

// connect to mongoDB
func connect(m config.Mongo) {
	c, err := mgo.M()

	if c == nil && err.Error() != "no connection" {
		logger.WithError(err).Panic("Connect Mongo failed")
	}

	if c != nil && err == nil {
		return
	}

	if err := mgo.Con(mgo.DBConfig{
		Host:       m.Host,
		User:       m.User,
		Password:   m.Password,
		Database:   m.Database,
		ReplicaSet: m.Replicaset,
	}); err != nil {
		logger.WithError(err).Panic("Connect Mongo failed")
	}
}

// Transaction -
func Transaction(fn func(sessCtx mongo.SessionContext) (interface{}, error)) error {
	client, err := mgo.M()
	if err != nil {
		logger.WithError(err).Error("connect mongo failed")
		return err
	}

	// second 10s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	session, err := client.StartSession()
	if err != nil {
		logger.WithError(err).Error("transaction start session failed")
		return err
	}

	defer session.EndSession(ctx)

	if _, err := session.WithTransaction(ctx, fn); err != nil {
		logger.WithError(err).Error("transaction session.WithTransaction failed")
		return err
	}

	return nil
}
