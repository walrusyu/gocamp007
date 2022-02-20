package configs

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
)

var (
	cfg  *TomlConfig
	once sync.Once
)

func GetConfig() *TomlConfig {
	once.Do(func() {
		filePath, err := filepath.Abs("./week4/configs/dev.toml")
		if err != nil {
			panic(err)
		}
		cfg = &TomlConfig{}
		if _, err := toml.DecodeFile(filePath, cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}

type TomlConfig struct {
	Server ServerConfig
}

type ServerConfig struct {
	Grpc GrpcConfig
}

type GrpcConfig struct {
	Port int
}
