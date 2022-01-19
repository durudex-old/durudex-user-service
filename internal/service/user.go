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
)

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
func (s *UserService) Create(ctx context.Context, user domain.User) (uint64, error) {
	// Validate user.
	if err := user.Validate(); err != nil {
		return 0, err
	}

	// Hasing user password.
	hashPassword, err := s.hash.Hash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashPassword

	// Creating a new user in postgres datatabe.
	id, err := s.repos.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
