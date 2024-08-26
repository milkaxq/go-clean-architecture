package configuration

import (
	"sync"
)

var (
	once   sync.Once
	config *Configs
)

type Configs struct {
	Server struct {
		Port int `mapstructure:"port"`
	}
	Database struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Dbname   string `mapstructure:"dbname"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
	}
}
