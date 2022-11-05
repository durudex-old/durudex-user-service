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
	"fmt"

	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/go-shared/database/postgres"

	"github.com/segmentio/ksuid"
)

// User repository interface.
type User interface {
	// Creating a new user in the database.
	Create(ctx context.Context, input domain.CreateUserInput) error

	// Getting a user from the database by his ID.
	Get(ctx context.Context, uid ksuid.KSUID) (domain.User, error)

	// Getting a user from the database using credentials."key" argument is the name of the
	// credential column in the database, "value" argument is the value of the credential.
	GetByCreds(ctx context.Context, key domain.CredentialKey, value string) (domain.User, error)
}

// User postgres repository structure.
type UserRepository struct{ psql postgres.Driver }

// Creating a new user postgres repository.
func NewUserRepository(psql postgres.Driver) User {
	return &UserRepository{psql: psql}
}

// Creating a new user in the database.
func (r *UserRepository) Create(ctx context.Context, input domain.CreateUserInput) error {
	query := `INSERT INTO users (id, username, email, password_hash, password_epoch)
		VALUES ($1, $2, $3, $4, $5)`

	// Inserting a new user in database.
	if _, err := r.psql.Exec(ctx, query, input.ID, input.Username, input.Email,
		input.PasswordHash, input.PasswordEpoch,
	); err != nil {
		return postgres.ErrorHandler(err, "User")
	}

	return nil
}

// Getting a user from the database by his ID.
func (r *UserRepository) Get(ctx context.Context, uid ksuid.KSUID) (domain.User, error) {
	query := "SELECT username, verified, avatar_url FROM users WHERE id=$1"

	// Selecting public user data using ID.
	row := r.psql.QueryRow(ctx, query, uid)

	var user domain.User

	// Scanning the received results from the database.
	if err := row.Scan(&user.Username, &user.Verified, &user.AvatarURL); err != nil {
		return domain.User{}, postgres.ErrorHandler(err, "User")
	}

	return user, nil
}

// Getting a user from the database using credentials."key" argument is the name of the credential column in the database,
// "value" argument is the value of the credential.
func (r *UserRepository) GetByCreds(ctx context.Context, key domain.CredentialKey, value string) (domain.User, error) {
	query := fmt.Sprintf(`SELECT id, username, email, password_hash, password_epoch, verified, avatar_url
		FROM users WHERE %s = $1`, key)

	// Select all user data using credentials.
	row := r.psql.QueryRow(ctx, query, value)

	var user domain.User

	// Scanning the received results from the database.
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.PasswordEpoch,
		&user.Verified, &user.AvatarURL,
	); err != nil {
		fmt.Println(err)
		return domain.User{}, postgres.ErrorHandler(err, "User")
	}

	return user, nil
}
