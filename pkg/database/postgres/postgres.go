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

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Postgres pool connections interface.
type Pool interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

// Postgres config.
type PostgresConfig struct {URL string}

// Create a new pool connection to postgres database.
func NewPostgresPool(cfg PostgresConfig) (*pgxpool.Pool, error) {
	// Connect to postgres database.
	pool, err := pgxpool.Connect(context.Background(), cfg.URL)
	if err != nil {
		return nil, errors.New("error pool connection to postgres database: " + err.Error())
	}

	// Check for connection operation.
	if err := pool.Ping(context.Background()); err != nil {
		return nil, errors.New("error connecting to postgres database connection: " + err.Error())
	}

	return pool, nil
}
