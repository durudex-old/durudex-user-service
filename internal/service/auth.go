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
	v1 "github.com/durudex/durudex-user-service/pkg/pb/durudex/v1"

	"github.com/segmentio/ksuid"
)

// Auth service interface.
type Auth interface {
	SignUp(ctx context.Context, user domain.User, code uint64, ip string) (domain.Tokens, error)
	SignIn(ctx context.Context, username, password, ip string) (domain.Tokens, error)
	SignOut(ctx context.Context, token, ip string) error
	RefreshTokens(ctx context.Context, token, ip string) (string, error)
	CreateSession(ctx context.Context, id ksuid.KSUID, ip string) (domain.Tokens, error)
}

// Auth service structure.
type AuthService struct {
	user    User
	code    Code
	email   v1.EmailUserServiceClient
	session postgres.Session
	cfg     *config.AuthConfig
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
	tokens, err := s.CreateSession(ctx, id, ip)
	if err != nil {
		return domain.Tokens{}, err
	}

	// Sending an email to a user with register.
	if _, err := s.email.SendEmailUserRegister(ctx, &v1.SendEmailUserRegisterRequest{
		Email:    user.Email,
		Username: user.Username,
	}); err != nil {
		return domain.Tokens{}, err
	}

	return tokens, nil
}

// User Sign In.
func (s *AuthService) SignIn(ctx context.Context, username, password, ip string) (domain.Tokens, error) {
	// Getting a user by credentials.
	user, err := s.user.GetByCreds(ctx, username, password)
	if err != nil {
		return domain.Tokens{}, err
	}

	// Creating a new user session.
	tokens, err := s.CreateSession(ctx, user.Id, ip)
	if err != nil {
		return domain.Tokens{}, err
	}

	// Sending an email to a user with logged in.
	if _, err := s.email.SendEmailUserLoggedIn(ctx, &v1.SendEmailUserLoggedInRequest{
		Email: user.Email,
		Ip:    ip,
	}); err != nil {
		return domain.Tokens{}, err
	}

	return tokens, nil
}

// User Sign Out.
func (s *AuthService) SignOut(ctx context.Context, token, ip string) error {
	// Deleting a user session by refresh token and ip address.
	return s.session.Delete(ctx, token, ip)
}

// Refresh user session access token.
func (s *AuthService) RefreshTokens(ctx context.Context, token, ip string) (string, error) {
	// Getting a user id in session by refresh token and ip address.
	id, err := s.session.GetUserId(ctx, token, ip)
	if err != nil {
		return "", err
	}

	// Generating a new jwt access token.
	accessToken, err := auth.GenerateAccessToken(id.String(), s.cfg.JWT.SigningKey, s.cfg.JWT.TTL)
	if err != nil {
		return "", err
	}

	return accessToken, nil
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
