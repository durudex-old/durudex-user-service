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

package server

import (
	"net"

	"github.com/durudex/durudex-user-service/internal/config"
	"github.com/durudex/durudex-user-service/internal/delivery/grpc"

	"github.com/rs/zerolog/log"
)

// The main structure of the server.
type Server struct {
	listener net.Listener
	grpc     *gRPCServer
	handler  *grpc.Handler
}

// Creating a new server.
func NewServer(cfg *config.ServerConfig, handler *grpc.Handler) (*Server, error) {
	log.Debug().Msg("Creating a new server...")

	// Creating a new TCP connections.
	lis, err := net.Listen("tcp", cfg.Host+":"+cfg.Port)
	if err != nil {
		return nil, err
	}

	// Creating a new gRPC server.
	grpcServer, err := NewGRPC(&cfg.TLS)
	if err != nil {
		return nil, err
	}

	return &Server{listener: lis, grpc: grpcServer, handler: handler}, nil
}

// Running server.
func (s *Server) Run() {
	log.Debug().Msg("Register gRPC handlers...")

	// Register gRPC handlers.
	s.handler.RegisterHandlers(s.grpc.Server)

	log.Info().Msg("Running server...")

	if err := s.grpc.Server.Serve(s.listener); err != nil {
		log.Fatal().Err(err).Msg("error running gRPC server")
	}
}

// Stopping server.
func (s *Server) Stop() {
	log.Info().Msg("Stoping gRPC server...")

	s.grpc.Server.Stop()
}
