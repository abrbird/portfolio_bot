package crypto_compare

type Historical struct {
	Response   string   `json:"Response"`
	Message    string   `json:"Message"`
	HasWarning bool     `json:"HasWarning"`
	Type       int      `json:"Type"`
	RateLimit  struct{} `json:"RateLimit"`
	Data       struct {
		Aggregated bool `json:"Aggregated"`
		TimeFrom   int  `json:"TimeFrom"`
		TimeTo     int  `json:"TimeTo"`
		Data       []struct {
			Time             int     `json:"time"`
			High             int     `json:"high"`
			Low              int     `json:"low"`
			Open             int     `json:"open"`
			Volumefrom       float64 `json:"volumefrom"`
			Volumeto         float64 `json:"volumeto"`
			Close            int     `json:"close"`
			ConversionType   string  `json:"conversionType"`
			ConversionSymbol string  `json:"conversionSymbol"`
		} `json:"Data"`
	} `json:"Data"`
}
