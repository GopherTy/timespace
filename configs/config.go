package configs

import (
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"

	"github.com/gopherty/timespace/pkg/xpath"

	"github.com/gopherty/timespace/pkg/xfile"

	"gopkg.in/yaml.v2"
)

var (
	conf *Config
)

// Config .
type Config struct {
	DB     DB     `json:"db" yaml:"DB"`
	Server Server `json:"server" yaml:"Server"`
}

// DB database config
type DB struct {
	// Driver database driven
	Driver string `json:"driver" yaml:"Driver"`
	// Source connection string
	Source string `json:"source" yaml:"Source"`
	// ShowSQL whether to display the SQL statement
	ShowSQL bool `json:"showSQL" yaml:"ShowSQL"`
	// MaxOpenConn number of database connections
	MaxOpenConn int `json:"maxOpenConn" yaml:"MaxOpenConn" `
	// MaxIdleConn maximum number of idle database connections
	MaxIdleConn int `json:"maxIdleConn" yaml:"MaxIdleConn" `
	// Cache cache size
	Cache int `json:"cache" yaml:"Cache"`
}

// Server server config
type Server struct {
	// Address listen address
	Address string `json:"address" yaml:"Address"`
	// CertFile certificate verification file
	CertFile string `json:"CertFile" yaml:"CertFile"`
	// KeyFile certificate
	KeyFile string `json:"keyFile" yaml:"KeyFile"`
	// Release is it a release version
	Release bool `json:"release" yaml:"Release"`
}

// load config
func (c *Config) loader(path string) (b []byte, err error) {
	path = strings.ToLower(strings.TrimSpace(path))
	switch filepath.Ext(path) {
	case ".json":
		b, err = xfile.ReadFile(path)
		if err != nil {
			return
		}

		err = json.Unmarshal(b, c)
		if err != nil {
			return
		}
	case ".yml":
		b, err = xfile.ReadFile(path)
		if err != nil {
			return
		}

		err = yaml.Unmarshal(b, c)
		if err != nil {
			return
		}
	default:
		err = errors.New("unknown format")
	}
	return
}

// Init load config file
func Init() (err error) {
	conf = &Config{}
	path, err := xpath.BasePath()
	if err != nil {
		return
	}

	path = filepath.Join(path, "config.yml")
	_, err = conf.loader(path)
	if err != nil {
		return
	}

	_, err = conf.loader(path)
	if err != nil {
		return
	}

	if conf == nil {
		panic("init failed. config is nil")
	}
	return
}

// Instance config object
func Instance() *Config {
	return conf
}
