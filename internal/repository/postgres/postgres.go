/*
 * Copyright © 2022 Durudex

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
	"github.com/durudex/durudex-user-service/internal/config"
	"github.com/durudex/durudex-user-service/pkg/database/postgres"

	"github.com/rs/zerolog/log"
)

// Postgres repository structure.
type PostgresRepository struct{ User }

// Creating a new postgres repository.
func NewPostgresRepository(cfg config.PostgresConfig) *PostgresRepository {
	log.Debug().Msg("Creating a new postgres repository")

	// Creating a new postgres pool connection.
	client, err := postgres.NewPool(&postgres.PostgresConfig{
		URL:      cfg.URL,
		MaxConns: cfg.MaxConns,
		MinConns: cfg.MinConns,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create postgres client")
	}

	return &PostgresRepository{User: NewUserRepository(client)}
}
