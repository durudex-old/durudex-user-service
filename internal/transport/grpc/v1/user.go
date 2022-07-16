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

	"github.com/durudex/dugopb/type/timestamp"
	"github.com/durudex/durudex-user-service/internal/service"
	v1 "github.com/durudex/durudex-user-service/pkg/pb/durudex/v1"

	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// User gRPC handler.
type UserHandler struct {
	service service.User
	v1.UnimplementedUserServiceServer
}

// Creating a new user gRPC handler.
func NewUserHandler(service service.User) *UserHandler {
	return &UserHandler{service: service}
}

// Getting user by id.
func (h *UserHandler) GetUserById(ctx context.Context, input *v1.GetUserByIdRequest) (*v1.GetUserByIdResponse, error) {
	// Getting user id from bytes.
	id, err := ksuid.FromBytes(input.Id)
	if err != nil {
		return &v1.GetUserByIdResponse{}, status.Error(codes.InvalidArgument, "Invalid Id")
	}

	// Getting user by id.
	user, err := h.service.GetByID(ctx, id)
	if err != nil {
		return &v1.GetUserByIdResponse{}, err
	}

	return &v1.GetUserByIdResponse{
		Username:  user.Username,
		LastVisit: timestamp.New(user.LastVisit),
		Verified:  user.Verified,
		AvatarUrl: user.AvatarUrl,
	}, nil
}

// Getting user by credentials.
func (h *UserHandler) GetUserByCreds(ctx context.Context, input *v1.GetUserByCredsRequest) (*v1.GetUserByCredsResponse, error) {
	// Getting user by credentials.
	user, err := h.service.GetByCreds(ctx, input.Username, input.Password)
	if err != nil {
		return &v1.GetUserByCredsResponse{}, err
	}

	return &v1.GetUserByCredsResponse{
		Id:        user.Id.Bytes(),
		Email:     user.Email,
		LastVisit: timestamp.New(user.LastVisit),
		Verified:  user.Verified,
		AvatarUrl: user.AvatarUrl,
	}, nil
}

// Forgot user password.
func (h *UserHandler) ForgotUserPassword(ctx context.Context, input *v1.ForgotUserPasswordRequest) (*v1.ForgotUserPasswordResponse, error) {
	// Forgot user password.
	if err := h.service.ForgotPassword(ctx, input.Password, input.Email, input.Code); err != nil {
		return &v1.ForgotUserPasswordResponse{}, err
	}

	return &v1.ForgotUserPasswordResponse{}, nil
}

func (h *UserHandler) UpdateUserAvatar(ctx context.Context, input *v1.UpdateUserAvatarRequest) (*v1.UpdateUserAvatarResponse, error) {
	return &v1.UpdateUserAvatarResponse{}, nil
}
