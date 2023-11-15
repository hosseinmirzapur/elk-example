package db

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog"
)

type Database struct {
	Conn   *sql.DB
	Logger zerolog.Logger
}

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
	Logger   zerolog.Logger
}

func Init(cfg Config) (*Database, error) {
	db := new(Database)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DbName)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.Conn = conn
	db.Logger = cfg.Logger

	err = db.Conn.Ping()

	return db, err
}
