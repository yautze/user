package consul

import (
	"log"

	"github.com/yautze/tools/config"
	"github.com/go-playground/validator"
	"github.com/hashicorp/consul/api"
)

var (
	Clinet *api.Client
)

// kvRepository -
type kvRepository struct {
	client  *api.Client
	options Options
}

// NewKVRepository -
func NewKVRepository(opts ...Option) *kvRepository {
	options := newOptions(opts...)

	k := &kvRepository{
		options: *options,
	}

	configure(k, options)

	return k
}

// configure -
func configure(k *kvRepository, opts *Options) {
	// connect consul agent
	if len(opts.Addrs) != 1 {
		log.Fatal("Addrs set failed")
	}

	c := &api.Config{
		Address: opts.Addrs[0],
	}

	client, err := api.NewClient(c)
	if err != nil {
		log.Fatal(err)
	}

	k.client = client

	Clinet = client
}

// Decode -
func (k *kvRepository) Decode(target interface{}) error {
	return setKV(k, target)
}

// setKV -
func setKV(k *kvRepository, target interface{}) error {
	// connection consul
	v, err := config.ConsulByClient(k.client, k.options.Name, string(k.options.ConfigType))
	if err != nil {
		log.Fatal(err)
		return err

	}

	// conv to struct
	err = v.Unmarshal(&target)
	if err != nil {
		log.Fatal(err)
		return err

	}

	// validate parameter
	validate := validator.New()
	if err := validate.Struct(target); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// GetClient -
func GetClient() *api.Client {
	return Clinet
}
