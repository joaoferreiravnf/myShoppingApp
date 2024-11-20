// config/config.go
package config

import (
	"database/sql"
	"fmt"
	"github.com/getsops/sops/v3/decrypt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	AppAuth        `yaml:"app_auth"`
	DatabaseConfig `yaml:"db"`
	GoogleAuth     `yaml:"google_auth"`
}

type AppAuth struct {
	Username string `yaml:"app_user"`
	Password string `yaml:"app_pass"`
}

type DatabaseConfig struct {
	Host     string `yaml:"db_host"`
	Port     int    `yaml:"db_port"`
	User     string `yaml:"db_user"`
	Password string `yaml:"db_password"`
	Name     string `yaml:"db_name"`
	Schema   string `yaml:"db_schema"`
	Table    string `yaml:"db_table"`
}

func LoadConfigs(filePath string) (*Config, error) {
	output, err := decrypt.File(filePath, "yaml")
	if err != nil {
		return nil, errors.Wrap(err, "failed to decrypt secrets with sops")
	}

	var config Config
	if err := yaml.Unmarshal(output, &config); err != nil {
		return nil, errors.Wrap(err, "failed to parse decrypted YAML")
	}

	return &config, nil
}

// ConnectToDatabase opens a new connection to the database
func ConnectToDatabase(config DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", config.Host, config.User, config.Password, config.Port, config.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &sql.DB{}, errors.Wrap(err, "error opening connection to db")
	}
	return db, nil
}

type GoogleAuth struct {
	ID     string
	Secret string
}
