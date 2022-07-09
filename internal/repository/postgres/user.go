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
	"errors"
	"fmt"

	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/pkg/database/postgres"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/segmentio/ksuid"
)

// User table name.
const UserTable string = "user"

// User repository interface.
type User interface {
	Create(ctx context.Context, user domain.User) (ksuid.KSUID, error)
	GetByID(ctx context.Context, id ksuid.KSUID) (domain.User, error)
	GetByUsername(ctx context.Context, username string) (domain.User, error)
	ForgotPassword(ctx context.Context, password, email string) error
	UpdateAvatar(ctx context.Context, avatarUrl string, id ksuid.KSUID) error
}

// User repository structure.
type UserRepository struct{ psql postgres.Postgres }

// Creating a new user repository.
func NewUserRepository(psql postgres.Postgres) *UserRepository {
	return &UserRepository{psql: psql}
}

// Creating a new user in postgres database.
func (r *UserRepository) Create(ctx context.Context, user domain.User) (ksuid.KSUID, error) {
	var id string

	// Query to create user.
	query := fmt.Sprintf(`INSERT INTO "%s" (username, email, password) VALUES ($1, $2, $3)
		RETURNING "id"`, UserTable)

	// Query and get user uuid.
	row := r.psql.QueryRow(ctx, query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		var pgErr *pgconn.PgError

		// Get postgres error.
		if errors.As(err, &pgErr) {
			// Switching postgres error code.
			if pgErr.Code == pgerrcode.UniqueViolation {
				// Return error if user with same username or email exists.
				return ksuid.Nil, &domain.Error{Code: domain.CodeAlreadyExists, Message: "User already exists"}
			}
		}

		return ksuid.Nil, &domain.Error{Code: domain.CodeInternal, Message: "Internal Server Error"}
	}

	return ksuid.Parse(id)
}

// Get user by id in postgres database.
func (r *UserRepository) GetByID(ctx context.Context, id ksuid.KSUID) (domain.User, error) {
	var user domain.User

	// Query for get user by id.
	query := fmt.Sprintf(`SELECT "username", "last_visit", "verified", "avatar_url"
		FROM "%s" WHERE "id"=$1`, UserTable)

	row := r.psql.QueryRow(ctx, query, id.String())

	// Scanning query row.
	err := row.Scan(&user.Username, &user.LastVisit, &user.Verified, &user.AvatarUrl)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, &domain.Error{Code: domain.CodeNotFound, Message: "User not found"}
		}
		fmt.Println(err)
		return domain.User{}, &domain.Error{Code: domain.CodeInternal, Message: "Internal Server Error"}
	}

	return user, nil
}

// Get user by username in postgres database.
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	var (
		user domain.User
		id   string
	)

	// Query for get user by username.
	query := fmt.Sprintf(`SELECT "id", "email", "password", "created_at", "last_visit", "verified",
	"avatar_url" FROM "%s" WHERE username=$1`, UserTable)

	row := r.psql.QueryRow(ctx, query, username)

	// Scanning query row.
	err := row.Scan(&id, &user.Email, &user.Password, &user.LastVisit,
		&user.Verified, &user.AvatarUrl)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, &domain.Error{Code: domain.CodeNotFound, Message: "User not found"}
		}

		return domain.User{}, &domain.Error{Code: domain.CodeInternal, Message: "Internal Server Error"}
	}

	// Parsing user id.
	user.Id, err = ksuid.Parse(id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// Forgot password in postgres database.
func (r *UserRepository) ForgotPassword(ctx context.Context, password, email string) error {
	// Query to update user password.
	query := fmt.Sprintf(`UPDATE "%s" SET password=$1 WHERE email=$2`, UserTable)
	_, err := r.psql.Exec(ctx, query, password, email)

	return err
}

// Update user avatar in postgres database.
func (r *UserRepository) UpdateAvatar(ctx context.Context, avatarUrl string, id ksuid.KSUID) error {
	// Query to update user avatar.
	query := fmt.Sprintf(`UPDATE "%s" SET "avatar_url"=$1 WHERE "id"=$2`, UserTable)
	_, err := r.psql.Exec(ctx, query, avatarUrl, id.String())

	return err
}
