package config

import "devtool/internal/repository"

type Config struct {
	DBConfig *repository.Config
}

func New() *Config {
	return &Config{
		DBConfig: &repository.Config{
			DriverName:     "sqlite3",
			DataSourceName: "devtool.db",
		},
	}
}