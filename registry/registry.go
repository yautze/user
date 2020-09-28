package registry

import "github.com/sarulabs/di"

// Container -
type Container struct {
	ctn di.Container
}

// New -
func New() (*Container, error) {
	container, err := newContainer(
		SetOptions("user-usecase", buidUserUsecase),
	)
	if err != nil {
		return nil, err
	}

	return container, nil
}

// newContainer -
func newContainer(opts ...Option) (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	options := &Options{}
	for _, o := range opts {
		o(options)
		d := di.Def{
			Name:  options.Name,
			Build: options.Build,
		}
		if err := builder.Add(d); err != nil {
			return nil, err
		}
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

// Resolve - get usecase by name
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

// Clean - clean all container
func (c *Container) Clean() error {
	return c.ctn.Clean()
}

// // New -
// func New() map[string]interface{} {
// 	ctn := make(map[string]interface{})

// 	ctn["user-usecase"] = bindUsecase("user")

// 	return ctn
// }

// func bindUsecase(u string) interface{} {
// 	userRepository := mongo.NewUserRepository(config.Config.Mongo)

// 	userService := service.NewUserService(userRepository)

// 	switch u {
// 	case "user":
// 		return usecase.NewUserUsecase(userRepository, userService)
// 	default:
// 		log.Panic("not usecase name")
// 	}

// 	return nil
// }
