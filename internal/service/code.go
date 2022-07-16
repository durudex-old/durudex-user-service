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

package service

import (
	"context"

	"github.com/durudex/durudex-user-service/internal/config"
	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/internal/repository/redis"
	"github.com/durudex/durudex-user-service/pkg/crypto/rand"
	v1 "github.com/durudex/durudex-user-service/pkg/pb/durudex/v1"
)

// Code service interface.
type Code interface {
	CreateVerifyEmailCode(ctx context.Context, email string) error
	VerifyEmailCode(ctx context.Context, email string, input uint64) (bool, error)
}

// Code service structure.
type CodeService struct {
	repos redis.Code
	email v1.EmailServiceClient
	cfg   *config.CodeConfig
}

// Creating a new code service.
func NewCodeService(repos redis.Code, email v1.EmailServiceClient, cfg *config.CodeConfig) *CodeService {
	return &CodeService{repos: repos, email: email, cfg: cfg}
}

// Creating a new user verification email code.
func (s *CodeService) CreateVerifyEmailCode(ctx context.Context, email string) error {
	// Generate random code.
	code, err := rand.Generate(s.cfg.MaxLength, s.cfg.MinLength)
	if err != nil {
		return err
	}

	// Creating a new code.
	if err := s.repos.CreateByEmail(ctx, email, code, s.cfg.TTL); err != nil {
		return err
	}

	// Sending an email to a user with a verification code.
	_, err = s.email.SendEmailUserCode(ctx, &v1.SendEmailUserCodeRequest{
		Email:    email,
		Username: "new user",
		Code:     code,
	})
	if err != nil {
		return err
	}

	return nil
}

// Verifying a user verification email code.
func (s *CodeService) VerifyEmailCode(ctx context.Context, email string, input uint64) (bool, error) {
	// Getting code by email.
	code, err := s.repos.GetByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	// Check input code.
	if input != code {
		return false, &domain.Error{Code: domain.CodeInvalidArgument, Message: "Invalid Code"}
	}

	return true, nil
}
