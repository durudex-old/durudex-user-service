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
	"github.com/durudex/durudex-user-service/pkg/database/postgres"

	"github.com/segmentio/ksuid"
)

// Session table name.
const SessionTable string = "user_session"

// User session repository interface.
type Session interface {
	Create(ctx context.Context, session domain.Session) error
	GetUserId(ctx context.Context, refreshToken, ip string) (ksuid.KSUID, error)
	Delete(ctx context.Context, refreshToken, ip string) error
}

// User session repository structure.
type SessionRepository struct{ psql postgres.Postgres }

// Creating a new user session repository.
func NewSessionRepository(psql postgres.Postgres) *SessionRepository {
	return &SessionRepository{psql: psql}
}

// Creating a new user session in postgres database.
func (r *SessionRepository) Create(ctx context.Context, session domain.Session) error {
	// Query to set a new user session in the postgres database.
	query := fmt.Sprintf(`INSERT INTO "%s" (id, user_id, refresh_token, ip, expires_in)
		VALUES ($1, $2, $3, $4, $5)`, SessionTable)
	_, err := r.psql.Exec(ctx, query, session.Id, session.UserId, session.RefreshToken,
		session.Ip, session.ExpiresIn)

	return err
}

// Getting user id by refresh token in postgres database.
func (r *SessionRepository) GetUserId(ctx context.Context, refreshToken, ip string) (ksuid.KSUID, error) {
	var id string

	// Query to get user id by refresh token from user session table.
	query := fmt.Sprintf(`SELECT user_id FROM "%s" WHERE refresh_token=$1 AND expires_in > now()
		AND ip=$2`, SessionTable)
	row := r.psql.QueryRow(ctx, query, refreshToken, ip)
	if err := row.Scan(&id); err != nil {
		return ksuid.Nil, err
	}

	return ksuid.Parse(id)
}

// Deleting a user session in postgres database.
func (r *SessionRepository) Delete(ctx context.Context, refreshToken, ip string) error {
	// Query to deleting user session by refresh token.
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE refresh_token=$1 AND ip=$2`, SessionTable)
	_, err := r.psql.Exec(ctx, query, refreshToken, ip)

	return err
}
