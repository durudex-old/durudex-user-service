/*
 * Copyright © 2021-2022 Durudex
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

package service

import (
	"context"

	"github.com/durudex/durudex-user-service/internal/config"
	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/internal/repository/postgres"
	"github.com/durudex/durudex-user-service/pkg/hash"

	"github.com/segmentio/ksuid"
)

// User service interface.
type User interface {
	Create(ctx context.Context, user domain.User) (ksuid.KSUID, error)
	GetByID(ctx context.Context, id ksuid.KSUID) (domain.User, error)
	GetByCreds(ctx context.Context, username, password string) (domain.User, error)
	ForgotPassword(ctx context.Context, password, email string, code uint64) error
	UpdateAvatar(ctx context.Context, id ksuid.KSUID, avatarUrl string) error
}

// User service structure.
type UserService struct {
	repos postgres.User
	code  Code
	cfg   *config.PasswordConfig
}

// Creating a new user service.
func NewUserService(repos postgres.User, code Code, cfg *config.PasswordConfig) *UserService {
	return &UserService{repos: repos, code: code, cfg: cfg}
}

// Creating a new user.
func (s *UserService) Create(ctx context.Context, user domain.User) (ksuid.KSUID, error) {
	var err error

	// Validate user.
	if err := user.Validate(); err != nil {
		return ksuid.Nil, err
	}

	// Generating a random user id.
	user.Id, err = ksuid.NewRandom()
	if err != nil {
		return ksuid.Nil, err
	}

	// Hashing user password.
	user.Password, err = hash.Hash(user.Password, s.cfg.Cost)
	if err != nil {
		return ksuid.Nil, err
	}

	// Creating a new user in postgres database.
	if err := s.repos.Create(ctx, user); err != nil {
		return ksuid.Nil, err
	}

	return user.Id, nil
}

// Getting user by id.
func (s *UserService) GetByID(ctx context.Context, id ksuid.KSUID) (domain.User, error) {
	// Getting user by id.
	user, err := s.repos.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// Getting user by credentials.
func (s *UserService) GetByCreds(ctx context.Context, username, password string) (domain.User, error) {
	// Getting user by username.
	user, err := s.repos.GetByUsername(ctx, username)
	if err != nil {
		return user, err
	}

	// Checking if user password is correct.
	if !hash.Check(user.Password, password) {
		return domain.User{}, &domain.Error{Code: domain.CodeInvalidArgument, Message: "Invalid Credentials"}
	}

	return user, nil
}

// Forgot user password.
func (s *UserService) ForgotPassword(ctx context.Context, password, email string, code uint64) error {
	// Verify email code.
	verify, err := s.code.VerifyEmailCode(ctx, email, code)
	if err != nil || !verify {
		return err
	}

	// Check user password.
	if !domain.RxPassword.MatchString(password) {
		return &domain.Error{Code: domain.CodeInvalidArgument, Message: "Invalid Password"}
	}

	// Hashing input user password.
	hashPassword, err := hash.Hash(password, s.cfg.Cost)
	if err != nil {
		return err
	}

	// Forgot password.
	return s.repos.ForgotPassword(ctx, hashPassword, email)
}

// Updating user avatar.
func (s *UserService) UpdateAvatar(ctx context.Context, id ksuid.KSUID, avatarUrl string) error {
	return s.repos.UpdateAvatar(ctx, avatarUrl, id)
}
