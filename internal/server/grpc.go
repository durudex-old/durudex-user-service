/*
	Copyright © 2022 Durudex

	This file is part of Durudex: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.

	Durudex is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with Durudex. If not, see <https://www.gnu.org/licenses/>.
*/

package server

import (
	"context"

	"github.com/durudex/durudex-user-service/internal/config"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	CACertFile          = "cert/rootCA.pem"
	userserviceCertFile = "cert/userservice-cert.pem"
	userserviceCertKey  = "cert/userservice-key.pem"
)

type GRPCServer struct {
	Server *grpc.Server
}

func NewGRPCServer(cfg *config.Config) *GRPCServer {
	serverOptions := []grpc.ServerOption{}

	// If TLS is true.
	if cfg.GRPC.TLS {
		tlsCredentials, err := LoadTLSCredentials(CACertFile, userserviceCertFile, userserviceCertKey)
		if err != nil {
			log.Fatal().Msgf("error load tls credentials: %s", err.Error())
		}

		serverOptions = append(
			serverOptions,
			grpc.Creds(tlsCredentials),
			grpc.UnaryInterceptor(unaryInterceptor),
		)
	}

	return &GRPCServer{Server: grpc.NewServer(serverOptions...)}
}

// Unary gtpc interceptor.
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Info().Msgf("unary interceptor: %s", info.FullMethod)

	return handler(ctx, req)
}
