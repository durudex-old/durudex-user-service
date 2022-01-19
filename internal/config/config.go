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
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type (
	// Config variables.
	Config struct {
		Server   ServerConfig   // Server config.
		Postgres PostgresConfig // Postgres config.
		Hash     HashConfig     // Hash config.
	}

	// Server config variables.
	ServerConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		TLS  bool   `mapstructure:"tls"`
	}

	// Hash config variables.
	HashConfig struct{ Password PasswordConfig }

	// Password config variables.
	PasswordConfig struct {
		Cost int `mapstructure:"cost"`
	}

	// Postgres config variables.
	PostgresConfig struct{ URL string }
)

// Initialize config.
func Init(configPath string) (*Config, error) {
	log.Debug().Msg("Initialize config...")

	// Populate defaults config variables.
	populateDefaults()

	// Parsing specified when starting the config file.
	if err := parseConfigFile(configPath); err != nil {
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
func parseConfigFile(configPath string) error {
	log.Debug().Msgf("Parsing config file: %s", configPath)

	// Split path to folder and file.
	path := strings.Split(configPath, "/")

	viper.AddConfigPath(path[0]) // Folder.
	viper.SetConfigName(path[1]) // File.

	// Read config file.
	return viper.ReadInConfig()
}

// Unmarshal config keys.
func unmarshal(cfg *Config) error {
	log.Debug().Msg("Unmarshal config keys...")

	// Unmarshal password keys.
	if err := viper.UnmarshalKey("hash.password", &cfg.Hash.Password); err != nil {
		return err
	}
	// Unmarshal server keys.
	return viper.UnmarshalKey("server", &cfg.Server)
}

// Set configurations from environment.
func setFromEnv(cfg *Config) {
	log.Debug().Msg("Set configurations from environment.")

	// Postgres configurations.
	cfg.Postgres.URL = os.Getenv("POSTGRES_URL")
}

// Populate defaults config variables.
func populateDefaults() {
	log.Debug().Msg("Populate defaults config variables...")

	// Server defaults.
	viper.SetDefault("server.host", defaultServerHost)
	viper.SetDefault("server.port", defaultServerPort)
	viper.SetDefault("server.tls", defaultServerTLS)

	// Password defaults.
	viper.SetDefault("hash.password.cost", defaultPasswordCost)
}
