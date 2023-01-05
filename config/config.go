package config

import (
	"devtool/internal/repository"
	"log"
	"os"
	"time"
)

type Config struct {
	DBConfig *repository.Config
}

func New(saveAfter string, restore bool, path string) *Config {
	if pathEnv, ok := os.LookupEnv("STORE_FILE"); ok {
		path = pathEnv
	}

	if restoreEnv, ok := os.LookupEnv("RESTORE"); ok {
		restore = true
		if restoreEnv == "false" {
			restore = false
		}
	}

	if saveAfterEnv, ok := os.LookupEnv("STORE_INTERVAL"); ok {
		saveAfter = saveAfterEnv
	}

	storeInterval, err := time.ParseDuration(saveAfter)
	if err != nil {
		log.Println(err)
		return nil
	}

	SaveAfter := storeInterval * time.Second

	return &Config{
		DBConfig: &repository.Config{
			DriverName:     "sqlite3",    // выбор между file, sqlite3, map
			DataSourceName: "devtool.db", // path
			SaveAfter:      SaveAfter,    // через сколько секунд изменения будут записываться
			Restore:        restore,      // восстанавливать ли предыдущие значения
		},
	}
}
