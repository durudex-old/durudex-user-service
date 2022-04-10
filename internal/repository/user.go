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

package repository

import (
	"context"

	"github.com/durudex/dugopg"
	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/internal/repository/psql"

	"github.com/gofrs/uuid"
)

// User repository interface.
type User interface {
	Create(ctx context.Context, user domain.User) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	GetByUsername(ctx context.Context, username string) (domain.User, error)
	ForgotPassword(ctx context.Context, password, email string) error
}

// User repository structure.
type UserRepository struct{ psql *psql.UserRepository }

// Creating a new user repository.
func NewUserRepository(conn dugopg.Native) *UserRepository {
	return &UserRepository{psql: psql.NewUserRepository(conn)}
}

// Creating a new user in datababe.
func (r *UserRepository) Create(ctx context.Context, user domain.User) (uuid.UUID, error) {
	return r.psql.Create(ctx, user)
}

// Get user by id in database.
func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return r.psql.GetByID(ctx, id)
}

// Get user by username in database.
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	return r.psql.GetByUsername(ctx, username)
}

// Forgot user password in database.
func (r *UserRepository) ForgotPassword(ctx context.Context, email, password string) error {
	return r.psql.ForgotPassword(ctx, password, email)
}
