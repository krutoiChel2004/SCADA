package database

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func newSQLServerDB(cfg Config) (*sqlx.DB, error) {
	// Формирование строки подключения
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s",
		cfg.Host, cfg.Username, cfg.Password, cfg.Port, cfg.DBName)

	// Открытие соединения
	db, err := sqlx.Open("sqlserver", dsn)
	if err != nil {
		return nil, err
	}

	// Проверка работоспособности соединения
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func GetDB() (*sqlx.DB, error) {
	db, err := newSQLServerDB(Config{
		Host:     "ite47\\DEVKP",
		Port:     "1438",
		Username: "ITE-NG_NT\\agafoma",
		Password: "ub43sz",
		DBName:   "compost",
	})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	return db, nil
}
