package configuration

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBConnection(config *Configs) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Dbname)

	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	if err = dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return dbpool, nil
}
