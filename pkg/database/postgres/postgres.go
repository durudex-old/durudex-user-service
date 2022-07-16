/*
 * Copyright © 2022 Durudex
 *
 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Postgres driver interface.
type Postgres interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

// Postgres config structure.
type PostgresConfig struct {
	URL      string
	MaxConns int32
	MinConns int32
}

// Configure postgres driver.
func (c *PostgresConfig) Configure(cfg *pgxpool.Config) {
	log.Debug().Msg("Configuring postgres driver")

	// Set driver logger.
	cfg.ConnConfig.Logger = zerologadapter.NewLogger(zerolog.New(os.Stderr))

	// Set max and min postgres driver connections.
	cfg.MaxConns = c.MaxConns
	cfg.MinConns = c.MinConns
}

// Creating a new postgres pool connection.
func NewPool(cfg *PostgresConfig) (Postgres, error) {
	log.Debug().Msg("Creating a new postgres pool connection")

	// Parsing database url.
	config, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, err
	}

	// Configure postgres driver.
	cfg.Configure(config)

	// Create a new pool connections by config.
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// Ping a database connection.
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
