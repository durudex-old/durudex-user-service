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

package config

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Default variables.
const (
	defaultGRPCHost = "userservice.durudex.local"
	defaultGRPCPort = "8004"
	defaultGRPCTLS  = true
)

type (
	Config struct {
		GRPC GRPCConfig
	}

	// GRPC server config variables.
	GRPCConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		TLS  bool   `mapstructure:"tls"`
	}
)

func Init(configPath string) (*Config, error) {
	log.Debug().Msg("Initialize config...")

	// Populate defaults config variables.
	populateDefaults()

	// Parsing config file.
	if err := parseConfigFile(configPath); err != nil {
		return nil, fmt.Errorf("error parsing config file: %s", err.Error())
	}

	var cfg Config
	// Unmarshal config keys.
	if err := unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshal config keys: %s", err.Error())
	}

	return &cfg, nil
}

// Parsing config file.
func parseConfigFile(configPath string) error {
	log.Debug().Msgf("Parsing config file: %s", configPath)

	// Split path to folder and file.
	path := strings.Split(configPath, "/")

	viper.AddConfigPath(path[0]) // folder
	viper.SetConfigName(path[1]) // file

	// Read config file.
	return viper.ReadInConfig()
}

// Unmarshal config keys.
func unmarshal(cfg *Config) error {
	log.Debug().Msg("Unmarshal config keys...")

	// Unmarshal grpc keys.
	return viper.UnmarshalKey("grpc", &cfg.GRPC)
}

// Populate defaults config variables.
func populateDefaults() {
	log.Debug().Msg("Populate defaults config variables...")

	viper.SetDefault("grpc.host", defaultGRPCHost)
	viper.SetDefault("grpc.port", defaultGRPCPort)
	viper.SetDefault("grpc.tls", defaultGRPCTLS)
}
