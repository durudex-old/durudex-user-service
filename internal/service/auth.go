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
	"time"

	"github.com/durudex/durudex-user-service/internal/config"
	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/internal/repository/postgres"
	"github.com/durudex/durudex-user-service/pkg/auth"

	"github.com/segmentio/ksuid"
)

// Auth service interface.
type Auth interface {
	SignUp(ctx context.Context, user domain.User, code uint64, ip string) (domain.Tokens, error)
	SignIn(ctx context.Context, username, password string) (domain.Tokens, error)
	SignOut(ctx context.Context, token, ip string) error
	RefreshTokens(ctx context.Context, token, ip string) (string, error)
	CreateSession(ctx context.Context, id ksuid.KSUID, ip string) (domain.Tokens, error)
}

// Auth service structure.
type AuthService struct {
	user    User
	code    Code
	session postgres.Session
	cfg     *config.AuthConfig
}

// Creating a new auth service.
func NewAuthService(user User, session postgres.Session, cfg *config.AuthConfig) *AuthService {
	return &AuthService{user: user, session: session, cfg: cfg}
}

// User Sign Up.
func (s *AuthService) SignUp(ctx context.Context, user domain.User, code uint64, ip string) (domain.Tokens, error) {
	// Verifying user email code.
	status, err := s.code.VerifyEmailCode(ctx, user.Email, code)
	if err != nil || !status {
		return domain.Tokens{}, err
	}

	// Creating a new user.
	id, err := s.user.Create(ctx, user)
	if err != nil {
		return domain.Tokens{}, err
	}

	// Creating a new user session.
	return s.CreateSession(ctx, id, ip)
}

func (s *AuthService) SignIn(ctx context.Context, username, password string) (domain.Tokens, error) {
	return domain.Tokens{}, nil
}

func (s *AuthService) SignOut(ctx context.Context, token, ip string) error {
	return nil
}

func (s *AuthService) RefreshTokens(ctx context.Context, token, ip string) (string, error) {
	return "", nil
}

// Creating a new user session.
func (s *AuthService) CreateSession(ctx context.Context, id ksuid.KSUID, ip string) (domain.Tokens, error) {
	// Generating a new jwt access token.
	accessToken, err := auth.GenerateAccessToken(id.String(), s.cfg.JWT.SigningKey, s.cfg.JWT.TTL)
	if err != nil {
		return domain.Tokens{}, err
	}

	// Generating a new refresh token.
	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		return domain.Tokens{}, err
	}

	// Creating a new user session.
	if err := s.session.Create(ctx, domain.Session{
		Id:           ksuid.New(),
		UserId:       id,
		RefreshToken: refreshToken,
		Ip:           ip,
		ExpiresIn:    time.Now().Add(s.cfg.Session.TTL),
	}); err != nil {
		return domain.Tokens{}, err
	}

	return domain.Tokens{Access: accessToken, Refresh: refreshToken}, nil
}
