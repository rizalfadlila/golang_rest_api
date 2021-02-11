package config

import (
	"github.com/pkg/errors"
	"github.com/rest_api/pkg/file"

	"sync"

	log "github.com/sirupsen/logrus"
)

var doOnce sync.Once
var singleton *Config

// NewConfig initialize config object
func NewConfig() *Config {
	doOnce.Do(func() {
		cfg, err := readCfg("./config.yaml")
		if err != nil {
			log.Fatalf(err.Error())
		}
		singleton = cfg
	})
	return singleton
}

// GetConfig :nodoc:
func GetConfig() *Config {
	if singleton != nil {
		return singleton
	}

	return &Config{
		App: &App{},
		Databases: &Databases{
			MongoDB: &MongoDB{},
		},
		Logger:      Logging{},
	}
}

// Config :nodoc:
type Config struct {
	App         *App       `yaml:"app"`
	Databases   *Databases `yaml:"databases"`
	Logger      Logging    `yaml:"logger"`
}

// App :nodoc:
type App struct {
	Env  string `yaml:"env"`
	Key  string `yaml:"key"`
	Port string `yaml:"port"`
}

// Databases :nodoc:
type Databases struct {
	MongoDB *MongoDB `yaml:"mongodb"`
}

// MongoDB :nodoc:
type MongoDB struct {
	Host           string              `yaml:"host"`
	Database       string              `yaml:"database"`
	URI            string              `yaml:"uri"`
	Authentication MongoAuthentication `yaml:"authentication"`
}

// MongoAuthentication :nodoc:
type MongoAuthentication struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// Logging :nodoc:
type Logging struct {
	SentryDSN string `yaml:"sentry_dsn"`
}


func readCfg(fname string) (*Config, error) {
	var cfg *Config

	err := file.ReadFromYAML(fname, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read file")
	}

	if cfg == nil {
		return nil, errors.New("No config file found")
	}

	return cfg, nil
}
