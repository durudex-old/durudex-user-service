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
	"context"

	"github.com/go-redis/redis/v8"
)

// Redis driver interface.
type Redis redis.Cmdable

// Creating a new redis client.
func NewClient(url string) (Redis, error) {
	// Parsing redis url.
	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	// Creating a new redis client.
	conn := redis.NewClient(opt)

	// Check for connections operation.
	if err := conn.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return conn, nil
}
