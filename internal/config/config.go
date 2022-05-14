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

package config

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Default config path.
const defaultConfigPath string = "configs/main"

type (
	// Config variables.
	Config struct {
		GRPC     GRPCConfig
		Database DatabaseConfig
		Hash     HashConfig
	}

	// gRPC server config variables.
	GRPCConfig struct {
		Host string    `mapstructure:"host"`
		Port string    `mapstructure:"port"`
		TLS  TLSConfig `mapstructure:"tls"`
	}

	// TLS config variables.
	TLSConfig struct {
		Enable bool   `mapstructure:"enable"`
		CACert string `mapstructure:"ca-cert"`
		Cert   string `mapstructure:"cert"`
		Key    string `mapstructure:"key"`
	}

	// Hash config variables.
	HashConfig struct {
		Password PasswordConfig `mapstructure:"password"`
	}

	// Password config variables.
	PasswordConfig struct {
		Cost int `mapstructure:"cost"`
	}

	// Database config variables.
	DatabaseConfig struct {
		Postgres PostgresConfig `mapstructure:"postgres"`
	}

	// Postgres config variables.
	PostgresConfig struct {
		MaxConns int32 `mapstructure:"max-conns"`
		MinConns int32 `mapstructure:"min-conns"`
		URL      string
	}
)

// Initialize config.
func Init() (*Config, error) {
	log.Debug().Msg("Initialize config...")

	// Parsing specified when starting the config file.
	if err := parseConfigFile(); err != nil {
		return nil, err
	}

	var cfg Config

	// Unmarshal config keys.
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Set configurations from environment.
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

// Unmarshal config keys.
func unmarshal(cfg *Config) error {
	log.Debug().Msg("Unmarshal config keys...")

	// Unmarshal password keys.
	if err := viper.UnmarshalKey("hash", &cfg.Hash); err != nil {
		return err
	}
	// Unmarshal postgres database keys.
	if err := viper.UnmarshalKey("database", &cfg.Database); err != nil {
		return err
	}
	// Unmarshal server keys.
	return viper.UnmarshalKey("grpc", &cfg.GRPC)
}

// Set configurations from environment.
func setFromEnv(cfg *Config) {
	log.Debug().Msg("Set configurations from environment.")

	// Postgres configurations.
	cfg.Database.Postgres.URL = os.Getenv("POSTGRES_URL")
}
