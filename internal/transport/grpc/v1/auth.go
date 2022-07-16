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

package v1

import (
	"context"

	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/internal/service"
	v1 "github.com/durudex/durudex-user-service/pkg/pb/durudex/v1"

	"github.com/segmentio/ksuid"
)

// User auth gRPC handler.
type AuthHandler struct {
	service service.Auth
	v1.UnimplementedUserAuthServiceServer
}

// Creating a new user auth gRPC handler.
func NewAuthHandler(service service.Auth) *AuthHandler {
	return &AuthHandler{service: service}
}

// User Sign Up gRPC handler.
func (h *AuthHandler) UserSignUp(ctx context.Context, input *v1.UserSignUpRequest) (*v1.UserSignUpResponse, error) {
	// User Sign Up.
	tokens, err := h.service.SignUp(ctx, domain.User{
		Id:       ksuid.New(),
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}, input.Code, input.Ip)
	if err != nil {
		return &v1.UserSignUpResponse{}, err
	}

	return &v1.UserSignUpResponse{Access: tokens.Access, Refresh: tokens.Refresh}, nil
}

// User Sign In gRPC handler.
func (h *AuthHandler) UserSignIn(ctx context.Context, input *v1.UserSignInRequest) (*v1.UserSignInResponse, error) {
	// User Sign In.
	tokens, err := h.service.SignIn(ctx, input.Username, input.Password, input.Ip)
	if err != nil {
		return &v1.UserSignInResponse{}, err
	}

	return &v1.UserSignInResponse{Access: tokens.Access, Refresh: tokens.Refresh}, nil
}

// User Sign Out gRPC handler.
func (h *AuthHandler) UserSignOut(ctx context.Context, input *v1.UserSignOutRequest) (*v1.UserSignOutResponse, error) {
	// User Sign Out.
	if err := h.service.SignOut(ctx, input.Refresh, input.Ip); err != nil {
		return &v1.UserSignOutResponse{}, err
	}

	return &v1.UserSignOutResponse{}, nil
}

// User Refresh token gRPC handler.
func (h *AuthHandler) RefreshUserToken(ctx context.Context, input *v1.RefreshUserTokenRequest) (*v1.RefreshUserTokenResponse, error) {
	// Refresh user token.
	access, err := h.service.RefreshTokens(ctx, input.Refresh, input.Ip)
	if err != nil {
		return &v1.RefreshUserTokenResponse{}, err
	}

	return &v1.RefreshUserTokenResponse{Access: access}, nil
}
