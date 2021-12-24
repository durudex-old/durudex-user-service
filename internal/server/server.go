/*
	Copyright © 2021 Durudex

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
	"net"

	"github.com/Durudex/durudex-user-service/internal/config"
	handler "github.com/Durudex/durudex-user-service/internal/delivery/grpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	CACertFile          = "cert/rootCA.pem"
	userserviceCertFile = "cert/userservice-cert.pem"
	userserviceCertKey  = "cert/userservice-key.pem"
)

type Server struct {
	tcpServer   *net.Listener
	grpcServer  *grpc.Server
	grpcHandler *handler.Handler
}

// Create a new server.
func NewServer(cfg *config.Config, grpcHandler *handler.Handler) *Server {
	serverOptions := []grpc.ServerOption{}

	// If TLS is true.
	if cfg.GRPC.TLS {
		tlsCredentials, err := handler.LoadTLSCredentials(CACertFile, userserviceCertFile, userserviceCertKey)
		if err != nil {
			log.Fatal().Msgf("error load tls credentials: %s", err.Error())
		}

		serverOptions = append(
			serverOptions,
			grpc.Creds(tlsCredentials),
			grpc.UnaryInterceptor(unaryInterceptor),
		)
	}

	// Server address.
	address := cfg.GRPC.Host + ":" + cfg.GRPC.Port

	return &Server{
		tcpServer:   NewTCPServer(address),
		grpcServer:  grpc.NewServer(serverOptions...),
		grpcHandler: grpcHandler,
	}
}

// Unary gtpc interceptor.
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Info().Msgf("unary interceptor: %s", info.FullMethod)
	return handler(ctx, req)
}

// Creating a new tcp server.
func NewTCPServer(address string) *net.Listener {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Msgf("error creating a new tcp serverL %s", err.Error())
		return nil
	}

	return &lis
}

// Run grpc server.
func (srv *Server) Run() {
	log.Debug().Msg("Running server...")

	// Registration services handlers.
	srv.grpcHandler.RegisterHandlers(srv.grpcServer)

	if err := srv.grpcServer.Serve(*srv.tcpServer); err != nil {
		log.Fatal().Msgf("error running grpc server: %s", err.Error())
	}
}

// Stop grpc server.
func (srv *Server) Stop() {
	log.Info().Msg("Stoping grpc server...")
	srv.grpcServer.Stop()
}
