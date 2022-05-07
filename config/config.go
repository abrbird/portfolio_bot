package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type File struct {
	Application     Application     `yaml:"application"`
	InternalAPIKeys InternalAPIKeys `yaml:"internalApiKeys"`
	ExternalAPIKeys ExternalAPIKeys `yaml:"externalApiKeys"`
	Database        Database        `yaml:"database"`
	DataSources     []DataSource    `yaml:"dataSources"`
}

type Application struct {
	ValidateInternal bool `yaml:"validateInternal"`
}

type InternalAPIKeys struct {
	AnyClient string `yaml:"anyClient"`
}

type ExternalAPIKeys struct {
	Telegram string `yaml:"telegram"`
}

type Database struct {
	DB       string `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type DataSource struct {
	Name   string `yaml:"name"`
	Url    string `yaml:"url"`
	ApiKey string `yaml:"apiKey"`
}

func (d *Database) Uri() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Password, d.DB)
}

func ParseFile(fileBytes []byte) (*File, error) {
	configFile := File{}

	err := yaml.Unmarshal(fileBytes, &configFile)
	if err != nil {
		return nil, err
	}

	return &configFile, nil
}

func ParseConfig(filepath string) (*File, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	configFile, err := ParseFile(b)
	if err != nil {
		return nil, err
	}

	return configFile, nil
}
