/*
 * Copyright © 2021-2022 Durudex
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

package repository

import (
	"github.com/durudex/durudex-user-service/internal/config"
	"github.com/durudex/durudex-user-service/internal/repository/postgres"
	"github.com/durudex/durudex-user-service/internal/repository/redis"
)

// Repository structure.
type Repository struct {
	Postgres *postgres.PostgresRepository
	Redis    *redis.RedisRepository
}

// Creating a new repository.
func NewRepository(config config.DatabaseConfig) *Repository {
	return &Repository{
		Postgres: postgres.NewPostgresRepository(config.Postgres),
		Redis:    redis.NewRedisRepository(config.Redis),
	}
}
