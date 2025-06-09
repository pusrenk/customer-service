package configs

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var (
	ConfigInstance *Config
	once           sync.Once
)

type Config struct {
	App      AppConfig      `json:"app" mapstructure:"app"`
	Database DatabaseConfig `json:"database" mapstructure:"database"`
}

type AppConfig struct {
	Name     string `json:"name" mapstructure:"name"`
	Version  string `json:"version" mapstructure:"version"`
	Port     int    `json:"port" mapstructure:"port"`
	GRPCPort int    `json:"grpc_port" mapstructure:"grpc_port"`
	Env      string `json:"env" mapstructure:"env"`
}

type DatabaseConfig struct {
	Host            string        `json:"host" mapstructure:"host"`
	Port            int           `json:"port" mapstructure:"port"`
	User            string        `json:"user" mapstructure:"user"`
	Password        string        `json:"password" mapstructure:"password"`
	Database        string        `json:"database" mapstructure:"database"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time" mapstructure:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime" mapstructure:"conn_max_lifetime"`
	MaxIdleConns    int           `json:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `json:"max_open_conns" mapstructure:"max_open_conns"`
}

// GetConfig get config instance
func GetConfig() *Config {
	return ConfigInstance
}

// InitConfig init config
func InitConfig() {
	once.Do(func() {
		ConfigInstance = &Config{}
		loadConfig(ConfigInstance)
	})
}

// loadConfig load config
func loadConfig(cfg *Config) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}
}
