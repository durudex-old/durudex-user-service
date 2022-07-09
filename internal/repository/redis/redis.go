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

package redis

import (
	"github.com/durudex/durudex-user-service/internal/config"
	"github.com/durudex/durudex-user-service/pkg/database/redis"

	"github.com/rs/zerolog/log"
)

// Redis repository structure.
type RedisRepository struct{ Code }

// Creating a new redis repository.
func NewRedisRepository(cfg config.RedisConfig) *RedisRepository {
	log.Debug().Msg("Creating a new redis repository")

	// Creating a new redis client.
	client, err := redis.NewClient(cfg.URL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create redis client")
	}

	return &RedisRepository{Code: NewCodeRepository(client)}
}
