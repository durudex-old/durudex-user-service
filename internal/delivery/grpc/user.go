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
	"github.com/durudex/durudex-user-service/internal/delivery/grpc/pb"
	"github.com/durudex/durudex-user-service/internal/service"
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
