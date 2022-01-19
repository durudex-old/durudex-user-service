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

package psql

import (
	"context"
	"errors"
	"fmt"

	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/pkg/database/postgres"
)

// User database tables.
const userTable = "user"

// User postgres repository structure.
type UserRepository struct{ psql postgres.Pool }

// Creating a new user postgres repository.
func NewUserRepository(psql postgres.Pool) *UserRepository {
	return &UserRepository{psql: psql}
}

// Creating a new user in postgres datatabe.
func (r *UserRepository) Create(ctx context.Context, user domain.User) (uint64, error) {
	var id uint64

	// Query to create user.
	query := fmt.Sprintf(`INSERT INTO "%s" (username, email, password) VALUES ($1, $2, $3) RETURNING "id"`, userTable)

	// Query and get id.
	row := r.psql.QueryRow(ctx, query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, errors.New("error creating a new user")
	}

	return id, nil
}
