package config

import "devtool/internal/repository"

type Config struct {
	DBConfig *repository.Config
}

func New() *Config {
	return &Config{
		DBConfig: &repository.Config{
			DriverName:     "map",      // выбор между file, sqlite3, map
			DataSourceName: "temp.txt", // path
		},
	}
}
