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

	v1 "github.com/durudex/durudex-user-service/pkg/pb/durudex/v1"
)

// User auth gRPC handler.
type AuthHandler struct {
	v1.UnimplementedUserAuthServiceServer
}

// Creating a new user auth gRPC handler.
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) UserSignUp(ctx context.Context, input *v1.UserSignUpRequest) (*v1.UserSignUpResponse, error) {
	return &v1.UserSignUpResponse{}, nil
}

func (h *AuthHandler) UserSignIn(ctx context.Context, input *v1.UserSignInRequest) (*v1.UserSignInResponse, error) {
	return &v1.UserSignInResponse{}, nil
}

func (h *AuthHandler) UserSignOut(ctx context.Context, input *v1.UserSignOutRequest) (*v1.UserSignOutResponse, error) {
	return &v1.UserSignOutResponse{}, nil
}

func (h *AuthHandler) RefreshUserToken(ctx context.Context, input *v1.RefreshUserTokenRequest) (*v1.RefreshUserTokenResponse, error) {
	return &v1.RefreshUserTokenResponse{}, nil
}
