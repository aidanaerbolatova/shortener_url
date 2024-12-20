package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"shortener-link/internal/config"
)

func NewConnection(config *config.Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		config.User,
		config.Password,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.SSLMode,
	)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
