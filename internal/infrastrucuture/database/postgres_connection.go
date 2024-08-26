package database

import (
	"booking/internal/domain/repositories"
	"booking/internal/infrastrucuture/configuration"
	"booking/internal/infrastrucuture/logging"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	pool    *pgxpool.Pool
	connStr string
	logger  *logging.ZapLogger
}

func NewDBConnection(config *configuration.Configs, logger *logging.ZapLogger) repositories.DBConnection {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Dbname)
	return &PostgresDB{connStr: connStr, logger: logger}
}

func (p *PostgresDB) Connect(ctx context.Context) error {
	var err error
	p.pool, err = pgxpool.New(ctx, p.connStr)
	if err != nil {
		p.logger.Error("Error in creating new pool", map[string]interface{}{})
		return err
	}

	p.logger.Info("Succesfully connected to pgxpool", map[string]interface{}{})

	if err = p.pool.Ping(ctx); err != nil {
		p.logger.Error("Error in ping to database", map[string]interface{}{})
		return err
	}

	p.logger.Info("Succesfully handled ping", map[string]interface{}{})

	return nil
}

func (p *PostgresDB) Close(ctx context.Context) error {
	p.pool.Close()

	p.logger.Info("Succesfully closed pool", map[string]interface{}{})

	return nil
}
