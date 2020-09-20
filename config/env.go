package config

import (
	"time"
)

// app name
const (
	APP = "DDDUser"
)

// env
var (
	// version
	VERSION string
	// build id
	BUILD string
	// commit id
	COMMIT string
)

// server
var (
	PORT    string
	ISUSEIP bool
)

// time
var (
	ZONE = time.FixedZone("", 8*60*60)
)

// consul
var (
	ConsulAdress string
	Config       Configuration
)
