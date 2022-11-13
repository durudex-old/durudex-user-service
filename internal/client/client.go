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

package client

import (
	"github.com/durudex/durudex-user-service/internal/config"
	v1 "github.com/durudex/durudex-user-service/pkg/pb/durudex/v1"

	"github.com/durudex/go-shared/transport/grpc"
)

// Client stores client implementations of services.
type Client struct {
	// Code stores the client implementation for the gRPC service.
	Code *grpc.Connection[v1.UserCodeServiceClient]
}

// New returns a new client.
func New(cfg config.ServiceConfig) *Client {
	return &Client{Code: grpc.Connect(v1.NewUserCodeServiceClient, cfg.Code)}
}
