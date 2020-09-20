package usecase

import (
	"testing"
	"user/app/domin/model"
	"user/app/domin/service"
	"user/app/interface/persistence/mongo"
	"user/config"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser(t *testing.T) {
	config.Setup()

	userRepository := mongo.NewUserRepository(config.Config.Mongo)

	userService := service.NewUserService(userRepository)

	userUsecase := NewUserUsecase(userRepository, userService)

	in := &model.User{
		ID:       primitive.NewObjectID(),
		FBID:     "YauTzFBID",
		Password: "111111",
		Info: model.Info{
			Name:  "Yautz",
			Phone: "0912345678",
		},
	}

	err := userUsecase.Create(in)

	assert.NoError(t, err)
}

func TestGetByID(t *testing.T) {
	config.Setup()

	userRepository := mongo.NewUserRepository(config.Config.Mongo)

	userService := service.NewUserService(userRepository)

	userUsecase := NewUserUsecase(userRepository, userService)

	id, _ := primitive.ObjectIDFromHex("5f66fd9ee71299269f949a21")

	user, err := userUsecase.GetByID(id)

	assert.NoError(t, err)
	spew.Dump(user)
}

func TestGetByFBIDAndPassword(t *testing.T) {
	config.Setup()

	userRepository := mongo.NewUserRepository(config.Config.Mongo)

	userService := service.NewUserService(userRepository)

	userUsecase := NewUserUsecase(userRepository, userService)

	user, err := userUsecase.GetByFBIDAndPassword("YauTzFBID", "111111")

	assert.NoError(t, err)

	spew.Dump(user)
}

func TestUpdateInfo(t *testing.T) {
	config.Setup()

	userRepository := mongo.NewUserRepository(config.Config.Mongo)

	userService := service.NewUserService(userRepository)

	userUsecase := NewUserUsecase(userRepository, userService)

	id, _ := primitive.ObjectIDFromHex("5f66fd9ee71299269f949a21")

	info := &model.Info{
		Name: "YauTzTest",
		Phone: "testPhoen",
	}

	err := userUsecase.UpdateInfo(id,info)

	assert.NoError(t, err)
}

func TestUpdatePassword(t *testing.T) {
	config.Setup()

	userRepository := mongo.NewUserRepository(config.Config.Mongo)

	userService := service.NewUserService(userRepository)

	userUsecase := NewUserUsecase(userRepository, userService)

	id, _ := primitive.ObjectIDFromHex("5f66fd9ee71299269f949a21")
	fbID := "YauTzFBID"
	password := "1111111"

	err := userUsecase.UpdatePassword(id, fbID, password)

	assert.NoError(t, err)
}

func TestDeleteByID(t *testing.T) {
	config.Setup()

	userRepository := mongo.NewUserRepository(config.Config.Mongo)

	userService := service.NewUserService(userRepository)

	userUsecase := NewUserUsecase(userRepository, userService)

	id, _ := primitive.ObjectIDFromHex("5f66fd9ee71299269f949a21")

	err := userUsecase.DeleteByID(id)

	assert.NoError(t, err)
}