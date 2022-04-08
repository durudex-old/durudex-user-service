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

package service

import (
	"context"

	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/internal/repository"
	"github.com/durudex/durudex-user-service/pkg/hash"

	"github.com/gofrs/uuid"
)

// User service interface.
type User interface {
	Create(ctx context.Context, user domain.User) (uuid.UUID, error)
	GetByCreds(ctx context.Context, username, password string) (domain.User, error)
	ForgotPassword(ctx context.Context, password, email string) (bool, error)
}

// User service structure.
type UserService struct {
	repos repository.User
	hash  hash.Password
}

// Creating a new user service.
func NewUserService(repos repository.User, hash hash.Password) *UserService {
	return &UserService{repos: repos, hash: hash}
}

// Creating a new user.
func (s *UserService) Create(ctx context.Context, user domain.User) (uuid.UUID, error) {
	// Validate user.
	if err := user.Validate(); err != nil {
		return uuid.UUID{}, err
	}

	// Hasing user password.
	hashPassword, err := s.hash.Hash(user.Password)
	if err != nil {
		return uuid.UUID{}, err
	}
	user.Password = hashPassword

	// Creating a new user in postgres database.
	id, err := s.repos.Create(ctx, user)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

// Getting user by credentials.
func (s *UserService) GetByCreds(ctx context.Context, username, password string) (domain.User, error) {
	// Hashing input user password.
	hashPassword, err := s.hash.Hash(password)
	if err != nil {
		return domain.User{}, err
	}

	// Getting user by credentials.
	user, err := s.repos.GetByCreds(ctx, username, hashPassword)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Forgot user password.
func (s *UserService) ForgotPassword(ctx context.Context, password, email string) (bool, error) {
	// Hashing input user password.
	hashPassword, err := s.hash.Hash(password)
	if err != nil {
		return false, err
	}

	// Forgot password.
	if err := s.repos.ForgotPassword(ctx, hashPassword, email); err != nil {
		return false, err
	}

	return true, nil
}
