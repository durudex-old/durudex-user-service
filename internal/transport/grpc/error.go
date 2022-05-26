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

package grpc

import (
	"errors"

	"github.com/durudex/durudex-user-service/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gRPC server error handler.
func errorHandler(err error) error {
	var e *domain.Error

	// Check if error is a domain.Error.
	if errors.As(err, &e) {
		switch e.Code {
		case domain.CodeNotFound:
			// Return gRPC error with status code not found.
			return status.Error(codes.NotFound, e.Message)
		case domain.CodeAlreadyExists:
			// Return gRPC error with status code already exists.
			return status.Error(codes.AlreadyExists, e.Message)
		case domain.CodeInvalidArgument:
			// Return gRPC error with status code invalid argument.
			return status.Error(codes.InvalidArgument, e.Message)
		}
	}

	return err
}
