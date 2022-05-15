package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Application     Application
	ClientAPIKeys   ClientAPIKeys
	ExternalAPIKeys ExternalAPIKeys
	Database        Database
	DataSourcesMap  map[string]DataSource
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

func ParseConfig(filepath string) (*Config, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	configFile, err := ParseFile(b)
	if err != nil {
		return nil, err
	}

	dsMap := make(map[string]DataSource, len(configFile.DataSources))
	for _, ds := range configFile.DataSources {
		dsMap[ds.Name] = ds
	}

	config_ := Config{
		Application:     configFile.Application,
		ClientAPIKeys:   configFile.ClientAPIKeys,
		ExternalAPIKeys: configFile.ExternalAPIKeys,
		Database:        configFile.Database,
		DataSourcesMap:  dsMap,
	}

	return &config_, nil
}
