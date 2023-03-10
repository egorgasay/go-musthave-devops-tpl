package config

import (
	"devtool/internal/repository"
	"flag"
	"log"
	"os"
	"time"
)

type Config struct {
	Host     string
	DBConfig *repository.Config
}

type F struct {
	host      *string
	path      *string
	saveAfter *string
	restore   *bool
}

var f F

func init() {
	f.host = flag.String("a", "localhost:8080", "-a=host")
	f.saveAfter = flag.String("i", "5m", "-i=Seconds")
	f.path = flag.String("f", "/tmp/devops-metrics-db.json", "-f=path")
	f.restore = flag.Bool("r", true, "-r=restore")
}

var True = true
var False = false

func New() *Config {
	flag.Parse()

	if pathEnv, ok := os.LookupEnv("STORE_FILE"); ok {
		f.path = &pathEnv
	}

	if restoreEnv, ok := os.LookupEnv("RESTORE"); ok {
		f.restore = &True
		if restoreEnv == "false" {
			f.restore = &False
		}
	}

	if saveAfterEnv, ok := os.LookupEnv("STORE_INTERVAL"); ok {
		f.saveAfter = &saveAfterEnv
	}

	storeInterval, err := time.ParseDuration(*f.saveAfter)
	if err != nil {
		log.Println(err)
		return nil
	}

	if addr, ok := os.LookupEnv("ADDRESS"); ok && addr != "" {
		f.host = &addr
	}

	return &Config{
		Host: *f.host,
		DBConfig: &repository.Config{
			DriverName:     "file",        // выбор между sqlite3, file
			DataSourceName: *f.path,       // путь до файла или данные бд
			SaveAfter:      storeInterval, // через сколько секунд изменения будут записываться
			Restore:        *f.restore,    // восстанавливать ли предыдущие значения
		},
	}
}
