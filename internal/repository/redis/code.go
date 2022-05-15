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

package redis

import (
	"context"

	"github.com/durudex/durudex-user-service/pkg/database/redis"
)

// Redis module name.
const EmailCodeModule string = "emailcode"

// Code repository interface.
type Code interface {
	CreateByEmail(ctx context.Context, email string, code uint64) error
	GetByEmail(ctx context.Context, email string) (uint64, error)
}

// Code repository structure.
type CodeRepository struct{ redis redis.Redis }

// Creating a new code repository.
func NewCodeRepository(redis redis.Redis) *CodeRepository {
	return &CodeRepository{redis: redis}
}

func (r *CodeRepository) CreateByEmail(ctx context.Context, email string, code uint64) error {
	return nil
}

func (r *CodeRepository) GetByEmail(ctx context.Context, email string) (uint64, error) {
	return 0, nil
}
