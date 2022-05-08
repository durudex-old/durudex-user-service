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

package grpc

import (
	"context"

	"github.com/durudex/dugopb/types/timestamp"
	"github.com/durudex/durudex-user-service/internal/delivery/grpc/pb"
	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/internal/service"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// User handler structure.
type UserHandler struct {
	service *service.Service
	pb.UnimplementedUserServiceServer
}

// Creating a new gRPC user handler.
func NewUserHandler(service *service.Service) *UserHandler {
	return &UserHandler{service: service}
}

// Create user handler.
func (h *UserHandler) CreateUser(ctx context.Context, input *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// Creating a new user.
	id, err := h.service.User.Create(ctx, domain.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return &pb.CreateUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateUserResponse{Id: id.Bytes()}, nil
}

// Getting user by ID handler.
func (h *UserHandler) GetUserByID(ctx context.Context, input *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	// Get user uuid from bytes.
	id, err := uuid.FromBytes(input.Id)
	if err != nil {
		return &pb.GetUserByIDResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// Getting user by ID.
	user, err := h.service.User.GetByID(ctx, id)
	if err != nil {
		return &pb.GetUserByIDResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetUserByIDResponse{
		Username:  user.Username,
		CreatedAt: timestamp.New(user.CreatedAt),
		LastVisit: timestamp.New(user.LastVisit),
		Verified:  user.Verified,
		AvatarUrl: user.AvatarURL,
	}, nil
}

// Getting user by credentials handler.
func (h *UserHandler) GetUserByCreds(ctx context.Context, input *pb.GetUserByCredsRequest) (*pb.GetUserByCredsResponse, error) {
	// Getting user by credentials.
	user, err := h.service.User.GetByCreds(ctx, input.Username, input.Password)
	if err != nil {
		return &pb.GetUserByCredsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetUserByCredsResponse{
		Id:        user.ID.Bytes(),
		Email:     user.Email,
		CreatedAt: timestamp.New(user.CreatedAt),
		LastVisit: timestamp.New(user.LastVisit),
		Verified:  user.Verified,
		AvatarUrl: user.AvatarURL,
	}, nil
}

// Forgot user password handler.
func (h *UserHandler) ForgotUserPassword(ctx context.Context, input *pb.ForgotUserPasswordRequest) (*pb.ForgotUserPasswordResponse, error) {
	err := h.service.User.ForgotPassword(ctx, input.Password, input.Email)
	if err != nil {
		return &pb.ForgotUserPasswordResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.ForgotUserPasswordResponse{}, nil
}

// Update user avatar handler.
func (h *UserHandler) UpdateUserAvatar(ctx context.Context, input *pb.UpdateUserAvatarRequest) (*pb.UpdateUserAvatarResponse, error) {
	// Get user uuid from bytes.
	id, err := uuid.FromBytes(input.Id)
	if err != nil {
		return &pb.UpdateUserAvatarResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// Update user avatar.
	if err := h.service.User.UpdateAvatar(ctx, id, input.AvatarUrl); err != nil {
		return &pb.UpdateUserAvatarResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateUserAvatarResponse{}, nil
}

// Creating user verification email code handler.
func (h *UserHandler) CreateVerifyUserEmailCode(ctx context.Context, input *pb.CreateVerifyUserEmailCodeRequest) (*pb.CreateVerifyUserEmailCodeResponse, error) {
	// TODO: check user email.

	// Create user verification email code.
	if err := h.service.Code.CreateVerifyEmailCode(ctx, input.Email); err != nil {
		return &pb.CreateVerifyUserEmailCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateVerifyUserEmailCodeResponse{}, nil
}

// Verify user email handler.
func (h *UserHandler) VerifyUserEmailCode(ctx context.Context, input *pb.VerifyUserEmailCodeRequest) (*pb.VerifyUserEmailCodeResponse, error) {
	// Verify user email code.
	verify, err := h.service.Code.VerifyEmailCode(ctx, input.Email, input.Code)
	if err != nil {
		return &pb.VerifyUserEmailCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.VerifyUserEmailCodeResponse{Status: verify}, nil
}
