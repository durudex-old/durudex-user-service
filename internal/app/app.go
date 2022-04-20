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

package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/durudex/dugopg/logger/dugopgzerolog"
	"github.com/durudex/dugopg/pool"
	"github.com/durudex/dugopg/types"
	"github.com/durudex/dugopg/types/uuid"
	"github.com/durudex/durudex-user-service/internal/config"
	"github.com/durudex/durudex-user-service/internal/delivery/grpc"
	"github.com/durudex/durudex-user-service/internal/repository"
	"github.com/durudex/durudex-user-service/internal/server"
	"github.com/durudex/durudex-user-service/internal/service"
	"github.com/durudex/durudex-user-service/pkg/hash"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// A function that running the application.
func Run() {
	// Set logger mode.
	if os.Getenv("DEBUG") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Initialize config.
	cfg, err := config.Init()
	if err != nil {
		log.Error().Err(err).Msg("error initialize config")
	}

	log.Debug().Msg("Creating connections to postgres database...")

	// Creating a new postgres connections.
	psql, err := pool.NewPool(pool.Config{
		URL:      cfg.Database.Postgres.URL,
		MaxConns: cfg.Database.Postgres.MaxConns,
		MinConns: cfg.Database.Postgres.MinConns,
		Logger:   dugopgzerolog.Init(),
		Types:    []types.DataType{uuid.UUID},
	})
	if err != nil {
		log.Error().Msgf("error connection to postgres database: %s", err.Error())
	}

	// Managers.
	hashManager := hash.NewHash(cfg.Hash.Password.Cost)

	// Repository, Service, Handler.
	repos := repository.NewRepository(psql)
	service := service.NewService(repos, hashManager)
	grpcHandler := grpc.NewHandler(service)

	// Create a new server.
	srv, err := server.NewServer(&cfg.Server, grpcHandler)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating a new server")
	}

	// Run server.
	go srv.Run()

	// Quit in application.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// Stoping server.
	srv.Stop()

	log.Info().Msg("Durudex User Service stoping!")
}
