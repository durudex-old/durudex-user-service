/*
 * Copyright © 2021-2022 Durudex
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

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/durudex/durudex-user-service/internal/config"
	"github.com/durudex/durudex-user-service/internal/repository"
	"github.com/durudex/durudex-user-service/internal/service"
	"github.com/durudex/durudex-user-service/internal/transport/grpc"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Initialize application.
func init() {
	// Set logger mode.
	if os.Getenv("DEBUG") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

// A function that running the application.
func main() {
	// Initialize config.
	cfg, err := config.Init()
	if err != nil {
		log.Error().Err(err).Msg("error initialize config")
	}

	// Creating a new gRPC client.
	client := grpc.NewClient(cfg.Service)
	// Creating a new repository.
	repos := repository.NewRepository(cfg.Database)
	// Creating a new service.
	service := service.NewService(repos, cfg, client.Email)
	// Creating a new gRPC handler.
	handler := grpc.NewHandler(service, cfg.Service)

	// Create a new server.
	srv := grpc.NewServer(cfg.GRPC, handler)

	// Run server.
	go srv.Run()

	// Quit in application.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// Stopping server.
	srv.Stop()

	log.Info().Msg("Durudex User Service stopping!")
}
