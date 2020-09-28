package registry

import "github.com/sarulabs/di"

// Option -
type Option func(*Options)

// Options -
type Options struct {
	Name  string
	Build func(ctn di.Container) (interface{}, error)
}

// SetOptions -
func SetOptions(name string, fn func(ctn di.Container) (interface{}, error)) Option {
	return func(o *Options) {
		o.Name = name
		o.Build = fn
	}
}
