/*
 * Copyright © 2022 Durudex

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

package v1

import (
	"context"

	"github.com/durudex/durudex-user-service/internal/service"
	v1 "github.com/durudex/durudex-user-service/pkg/pb/durudex/v1"
)

// User gRPC server handler.
type UserHandler struct {
	service service.User
	v1.UnimplementedUserServiceServer
}

// Creating a new user gRPC handler.
func NewUserHandler(service service.User) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(ctx context.Context, input *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	return &v1.CreateUserResponse{}, nil
}

func (h *UserHandler) GetUserById(ctx context.Context, input *v1.GetUserByIdRequest) (*v1.GetUserByIdResponse, error) {
	return &v1.GetUserByIdResponse{}, nil
}

func (h *UserHandler) GetUserByCreds(ctx context.Context, input *v1.GetUserByCredsRequest) (*v1.GetUserByCredsResponse, error) {
	return &v1.GetUserByCredsResponse{}, nil
}

func (h *UserHandler) ForgotUserPassword(ctx context.Context, input *v1.ForgotUserPasswordRequest) (*v1.ForgotUserPasswordResponse, error) {
	return &v1.ForgotUserPasswordResponse{}, nil
}

func (h *UserHandler) UpdateUserAvatar(ctx context.Context, input *v1.UpdateUserAvatarRequest) (*v1.UpdateUserAvatarResponse, error) {
	return &v1.UpdateUserAvatarResponse{}, nil
}

func (h *UserHandler) CreateVerifyUserEmailCode(ctx context.Context, input *v1.CreateVerifyUserEmailCodeRequest) (*v1.CreateVerifyUserEmailCodeResponse, error) {
	return &v1.CreateVerifyUserEmailCodeResponse{}, nil
}

func (h *UserHandler) VerifyUserEmailCode(ctx context.Context, input *v1.VerifyUserEmailCodeRequest) (*v1.VerifyUserEmailCodeResponse, error) {
	return &v1.VerifyUserEmailCodeResponse{}, nil
}
