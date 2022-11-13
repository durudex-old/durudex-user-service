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

package config

import (
	"os"
	"path/filepath"

	"github.com/durudex/go-shared/transport/grpc"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Default config path.
const defaultConfigPath string = "configs/main"

type (
	// Config stores all configuration structures.
	Config struct {
		// GRPC server config variables.
		GRPC grpc.ServerConfig `mapstructure:"grpc"`

		// Database config variables.
		Database DatabaseConfig `mapstructure:"database"`

		// Service config variables.
		Service ServiceConfig `mapstructure:"service"`
	}

	// Database config variables.
	DatabaseConfig struct {
		// Postgres database config variables.
		Postgres PostgresConfig `mapstructure:"postgres"`
	}

	// Postgres config variables.
	PostgresConfig struct {
		// MaxConns stores the maximum number of database connections.
		MaxConns int32 `mapstructure:"max-conns"`

		// MinConns stores the minimum number of database connections.
		MinConns int32 `mapstructure:"min-conns"`

		// URL is used to establish a connection to the database, and it may also contain some
		// connection configuration.
		URL string
	}

	// Service config variables.
	ServiceConfig struct {
		// Code service gRPC client.
		Code grpc.ConnectionConfig `mapstructure:"code"`
	}
)

// New returns a new config.
func New() (*Config, error) {
	log.Debug().Msg("Creating a new config...")

	// Parsing specified when starting the config file.
	if err := parseConfigFile(); err != nil {
		return nil, err
	}

	var cfg Config

	// Unmarshal config keys.
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Sets configurations from environment.
	setFromEnv(&cfg)

	return &cfg, nil
}

// Parsing specified when starting the config file.
func parseConfigFile() error {
	// Get config path variable.
	configPath := os.Getenv("CONFIG_PATH")

	// Check is config path variable empty.
	if configPath == "" {
		configPath = defaultConfigPath
	}

	log.Debug().Msgf("Parsing config file: %s", configPath)

	// Split path to folder and file.
	dir, file := filepath.Split(configPath)

	viper.AddConfigPath(dir)
	viper.SetConfigName(file)

	// Read config file.
	return viper.ReadInConfig()
}

// Sets configurations from environment.
func setFromEnv(cfg *Config) {
	log.Debug().Msg("Set configurations from environment.")

	// Postgres database configurations.
	cfg.Database.Postgres.URL = os.Getenv("POSTGRES_URL")
}
