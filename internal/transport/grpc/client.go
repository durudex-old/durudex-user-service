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

package grpc

import (
	"github.com/durudex/durudex-user-service/internal/config"
	v1 "github.com/durudex/durudex-user-service/pkg/pb/durudex/v1"
	"github.com/durudex/durudex-user-service/pkg/tls"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// gRPC client structure.
type Client struct{ Email v1.EmailServiceClient }

// Creating a new gRPC client.
func NewClient(cfg config.ServiceConfig) *Client {
	return &Client{Email: v1.NewEmailServiceClient(connectToService(cfg.Email))}
}

// Connection to the gRPC server.
func connectToService(cfg config.Service) *grpc.ClientConn {
	log.Info().Msgf("Connecting to %s service", cfg.Addr)

	var opts []grpc.DialOption

	if cfg.TLS.Enable {
		creds, err := tls.LoadTLSConfig(cfg.TLS.CACert, cfg.TLS.Cert, cfg.TLS.Key)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to load TLS credentials")
		}

		// Append client credential options.
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(creds)))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// Creating a new gRPC client connection.
	conn, err := grpc.Dial(cfg.Addr, opts...)
	if err != nil {
		log.Error().Err(err).Msgf("failed to connect to service: %s", cfg.Addr)
	}

	return conn
}
