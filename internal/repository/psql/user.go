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
	"fmt"

	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/pkg/database/postgres"
)

// User database tables.
const userTable string = "user"

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
	query := fmt.Sprintf(`INSERT INTO "%s" (username, email, password) VALUES ($1, $2, $3)
		RETURNING "id"`, userTable)

	// Query and get id.
	row := r.psql.QueryRow(ctx, query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// Get user by credentials for postgres database.
func (r *UserRepository) GetByCreds(ctx context.Context, username, password string) (domain.User, error) {
	var user domain.User

	// Query and get user.
	query := fmt.Sprintf(`SELECT "id", "username", "email", "joined_in", "last_visit",
		"verified", "avatar_url" FROM "%s" WHERE username=$1 AND password=$2`, userTable)

	row := r.psql.QueryRow(ctx, query, username, password)

	// Scanning query row.
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.JoinedIn, &user.LastVisit,
		&user.Verified, &user.AvatarURL)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
