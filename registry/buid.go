package registry

import (
	"user/app/domin/service"
	"user/app/interface/persistence/mongo"
	"user/app/usecase"
	"user/config"

	"github.com/sarulabs/di"
)

// buidUserUsecase - buid auth usecase
func buidUserUsecase(ctn di.Container) (interface{}, error) {
	userRepository := mongo.NewUserRepository(config.Config.Mongo)
	userService := service.NewUserService(userRepository)

	return usecase.NewUserUsecase(userRepository, userService), nil
}
