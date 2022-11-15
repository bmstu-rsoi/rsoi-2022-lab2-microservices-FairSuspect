package apiserver

import (
	"http-rest-api/store"
)

// Config ...
type Config struct {
	BindAddr string `toml:"bond_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	// port := os.Getenv("PORT")
	// if len(port) == 0 {
	port := "8070"
	// }
	return &Config{
		BindAddr: ":" + port,
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
