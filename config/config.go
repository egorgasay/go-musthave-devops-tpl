package config

import (
	"devtool/internal/repository"
	"os"
	"strconv"
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

	saveAfterInt, err := strconv.Atoi(saveAfter)
	if err != nil {
		saveAfterInt = 300
	}

	saveAfterSeconds := time.Duration(saveAfterInt) * time.Second

	return &Config{
		DBConfig: &repository.Config{
			DriverName:     "file",           // выбор между file, sqlite3, map
			DataSourceName: path,             // path
			SaveAfter:      saveAfterSeconds, // через сколько секунд изменения будут записываться
			Restore:        restore,          // восстанавливать ли предыдущие значения
		},
	}
}
