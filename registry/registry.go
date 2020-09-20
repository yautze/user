package registry

import (
	"log"
	"user/app/domin/service"
	"user/app/interface/persistence/mongo"
	"user/app/usecase"
	"user/config"
)

// New -
func New() map[string]interface{} {
	ctn := make(map[string]interface{})

	ctn["user-usecase"] = bindUsecase("user")

	return ctn
}

func bindUsecase(u string) interface{} {
	userRepository := mongo.NewUserRepository(config.Config.Mongo)

	userService := service.NewUserService(userRepository)

	switch u {
	case "user":
		return usecase.NewUserUsecase(userRepository, userService)
	default:
		log.Panic("not usecase name")
	}

	return nil
}
