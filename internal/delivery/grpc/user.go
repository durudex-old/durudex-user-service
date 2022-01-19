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

	"github.com/durudex/durudex-user-service/internal/delivery/grpc/pb"
	"github.com/durudex/durudex-user-service/internal/delivery/grpc/pb/types"
	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/internal/service"

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
func (h *UserHandler) Create(ctx context.Context, input *pb.CreateRequest) (*types.ID, error) {
	// Create user model
	user := domain.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	// Creating a new user.
	id, err := h.service.Create(ctx, user)
	if err != nil {
		return &types.ID{Id: 0}, status.Error(codes.Internal, err.Error())
	}

	return &types.ID{Id: id}, nil
}
