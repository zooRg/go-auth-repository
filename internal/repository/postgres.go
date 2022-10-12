package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config, logger *zap.SugaredLogger) (*sqlx.DB, error) {
	dbUrl := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.DBName,
		cfg.Password,
		cfg.SSLMode,
	)
	logger.Debugf("db url: %s", dbUrl)

	db, err := sqlx.Open("postgres", dbUrl)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
