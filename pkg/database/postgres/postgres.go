/*
 * Copyright © 2021-2022 Durudex

 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.

 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package postgres

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

// Postgres pool connections interface.
type Pool interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

// Postgres config.
type PostgresConfig struct {
	MaxConns int32
	MinConns int32
	URL      string
}

// Create a new pool connection to postgres database.
func NewPostgresPool(cfg PostgresConfig) (*pgxpool.Pool, error) {
	// Parsing postgres config.
	config, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, errors.New("error parsing config: " + err.Error())
	}

	// Set max and min client pool connections.
	config.MaxConns = cfg.MaxConns
	config.MinConns = cfg.MinConns

	// Set postgres logger.
	config.ConnConfig.Logger = zerologadapter.NewLogger(zerolog.New(os.Stderr))

	// Connect to postgres database.
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, errors.New("error pool connection to postgres database: " + err.Error())
	}

	// Check for connection operation.
	if err := pool.Ping(context.Background()); err != nil {
		return nil, errors.New("error connecting to postgres database connection: " + err.Error())
	}

	return pool, nil
}
