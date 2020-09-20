package config

import (
	"sync"
	"user/app/interface/persistence/consul"

	"github.com/sirupsen/logrus"
	"github.com/yautze/tools/logger"
)

var once sync.Once

// Configuration -
type Configuration struct {
	Mongo Mongo `yaml:"mongo"`
	Log   Log   `yaml:"log"`
}

// Mongo -
type Mongo struct {
	Host       string `yaml:"host"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Database   string `yaml:"database"`
	Replicaset string `yaml:"replicaset"`
}

// Log -
type Log struct {
	Level string `yaml:"level"`
}

// Setup -
func Setup() {
	once.Do(func() {
		// connect consul and set KV
		consul.NewKVRepository(
			consul.SetName(APP),
			consul.SetAddrs(ConsulAdress),
		).Decode(&Config)

		// set Logger
		setLogger(Config.Log.Level)
	})
}

// set logger
func setLogger(level string) {
	logger.NewLogrus()

	logger.SetServiceInfo(APP)

	var l logrus.Level

	switch level {
	case "trace":
		l = logger.TraceLevel
	case "debug":
		l = logger.DebugLevel
	case "Info":
		l = logger.InfoLevel
	case "Error":
		l = logger.ErrorLevel
	case "Warn":
		l = logger.WarnLevel
	default:
		l = logger.DebugLevel
	}

	logger.SetLevel(l)
}
