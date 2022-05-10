package config

type File struct {
	Application     Application     `yaml:"application"`
	ClientAPIKeys   ClientAPIKeys   `yaml:"clientApiKeys"`
	ExternalAPIKeys ExternalAPIKeys `yaml:"externalApiKeys"`
	Database        Database        `yaml:"database"`
	DataSources     []DataSource    `yaml:"dataSources"`
}

type Application struct {
	HistoryStartTimeStamp int64        `yaml:"historyStartTimeStamp"`
	HistoryInterval       uint64       `yaml:"historyInterval"`
	ValidateInternal      bool         `yaml:"validateInternal"`
	BaseCurrency          string       `yaml:"baseCurrency"`
	AvailableMarketItems  []MarketItem `yaml:"availableMarketItems"`
}

type MarketItem struct {
	Code string `yaml:"code"`
	Type string `yaml:"type"`
}

type ClientAPIKeys struct {
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
