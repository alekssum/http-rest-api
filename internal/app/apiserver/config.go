package apiserver

// Config ...
type Config struct {
	BindAddr    string `toml:"bund_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"sessions_key"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
