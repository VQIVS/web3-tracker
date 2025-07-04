package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	GolbalConfig *Config
	once         sync.Once
)

type Config struct {
	Ethereum EthereumConfig `yaml:"ethereum"`
	Database DatabaseConfig `yaml:"database"`
	Server   ServerConfig   `yaml:"server"`
	Wallets  WalletsConfig  `yaml:"wallets"`
}

type EthereumConfig struct {
	RPCURL string `yaml:"rpc_url"`
}

type DatabaseConfig struct {
	Path string `yaml:"path"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type WalletsConfig struct {
	FilePath string `yaml:"file_path"`
}

func initConfig(fileName string) error {
	var err error
	once.Do(func() {
		GolbalConfig, err = LoadConfig(fileName)
	})
	return err
}

func MustInit(fileName string) {
	if err := initConfig(fileName); err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
}

func Get() Config {
	if GolbalConfig == nil {
		panic("config not initialized, call initConfig() or MustInit() first")
	}
	return *GolbalConfig
}

func LoadConfig(fileName string) (*Config, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &config, nil
}

func MustLoad(fileName string) Config {
	config, err := LoadConfig(fileName)
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	return *config
}
