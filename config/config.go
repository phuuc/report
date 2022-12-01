package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Environment string  `yaml:"environment"`
	Server      *Server `yaml:"server"`
	DB          *DB     `yaml:"db"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DB struct {
	DriverName                 string `yaml:"driver_name"`
	MaxOpenConns               int    `yaml:"max_open_conns"`
	MaxIdleConns               int    `yaml:"max_idle_conns"`
	DataSource                 string `yaml:"data_source"`
	MigrationConnURL           string `yaml:"migration_conn_url"`
	ConnMaxLifeTimeMiliseconds int64  `yaml:"conn_max_life_time_ms"`
}

func Load(filePath string) (*Config, error) {
	if len(filePath) == 0 {
		filePath = RootDir() + os.Getenv("CONFIG_FILE")
	}

	configBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	configBytes = []byte(os.ExpandEnv(string(configBytes)))

	cfg := &Config{}

	err = yaml.Unmarshal(configBytes, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(2)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
func Addr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
