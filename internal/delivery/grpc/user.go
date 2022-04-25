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
	service service.User
	pb.UnimplementedUserServiceServer
}

// Creating a new gRPC user handler.
func NewUserHandler(service service.User) *UserHandler {
	return &UserHandler{service: service}
}

// Create user handler.
func (h *UserHandler) CreateUser(ctx context.Context, input *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// Create user model
	user := domain.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	// Creating a new user.
	id, err := h.service.Create(ctx, user)
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
		return &pb.GetUserByIDResponse{}, err
	}

	// Getting user by ID.
	user, err := h.service.GetByID(ctx, id)
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
	user, err := h.service.GetByCreds(ctx, input.Username, input.Password)
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
func (h *UserHandler) ForgotPassword(ctx context.Context, input *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	err := h.service.ForgotPassword(ctx, input.Password, input.Email)
	if err != nil {
		return &pb.ForgotPasswordResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.ForgotPasswordResponse{}, nil
}
