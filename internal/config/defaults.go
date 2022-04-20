/*
 * Copyright © 2022 Durudex

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
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	// Config defaults.
	defaultConfigPath string = "configs/main"

	// Server defaults.
	defaultServerHost string = "user.service.durudex.local"
	defaultServerPort string = "8004"

	// TLS server defaults.
	defaultTLSEnable bool   = true
	defaultTLSCACert string = "./certs/rootCA.pem"
	defaultTLSCert   string = "./certs/user.service.durudex.local-cert.pem"
	defaultTLSKey    string = "./certs/user.service.durudex.local-key.pem"

	// Password defaults.
	defaultPasswordCost int = 14

	// Postgres defaults.
	defaultPostgresMaxConns int32 = 20
	defaultPostgresMinConns int32 = 5
)

// Populate defaults config variables.
func populateDefaults() {
	log.Debug().Msg("Populate defaults config variables...")

	// Server defaults.
	viper.SetDefault("server.host", defaultServerHost)
	viper.SetDefault("server.port", defaultServerPort)

	// TLS server defaults.
	viper.SetDefault("server.tls.enable", defaultTLSEnable)
	viper.SetDefault("server.tls.ca-cert", defaultTLSCACert)
	viper.SetDefault("server.tls.cert", defaultTLSCert)
	viper.SetDefault("server.tls.key", defaultTLSKey)

	// Password defaults.
	viper.SetDefault("hash.password.cost", defaultPasswordCost)

	// Postgres defaults.
	viper.SetDefault("database.postgres.maxConn", defaultPostgresMaxConns)
	viper.SetDefault("database.postgres.minConn", defaultPostgresMinConns)
}
