package consul

import (
	"log"
	"time"
)

var (
	DefaultTimeout    = 5 * time.Second
	DefaultAddrs      = []string{"127.0.0.1:8500"}
	DefaultConfigType = ConfigType_YAML
)

// Option -
type Option func(*Options)

// ConfigType -
type ConfigType string

const (
	ConfigType_YAML ConfigType = "yaml"
	ConfigType_JSON ConfigType = "json"
)

// Options -
type Options struct {
	Addrs      []string
	Name       string
	Timeout    time.Duration
	ConfigType ConfigType
}

// newOptions -
func newOptions(opts ...Option) *Options {
	options := &Options{
		Addrs:      DefaultAddrs,
		Timeout:    DefaultTimeout,
		ConfigType: DefaultConfigType,
	}

	for _, o := range opts {
		o(options)
	}

	if options.Name == "" {
		log.Fatal("not set name")
		return nil
	}

	return options
}

// SetAddrs - Addrs is the registry addresses to use
func SetAddrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

// SetName -
func SetName(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}