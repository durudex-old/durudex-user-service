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

package config_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/durudex/durudex-user-service/internal/config"

	"github.com/durudex/go-shared/crypto/tls"
	"github.com/durudex/go-shared/transport/grpc"
)

// Testing creating a new config.
func TestConfig_New(t *testing.T) {
	// Environment configurations.
	type env struct{ configPath, postgresURL, redisURL, jwtSigningKey string }

	// Testing args.
	type args struct{ env env }

	// Set environments configurations.
	setEnv := func(env env) {
		os.Setenv("CONFIG_PATH", env.configPath)
		os.Setenv("POSTGRES_URL", env.postgresURL)
	}

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		want    *config.Config
		wantErr bool
	}{
		{
			name: "OK",
			args: args{env: env{
				configPath:    "fixtures/main",
				postgresURL:   "postgres://localhost:1",
				redisURL:      "redis://user.redis.durudex.local:6379",
				jwtSigningKey: "secret-key",
			}},
			want: &config.Config{
				GRPC: grpc.ServerConfig{
					Host: "user.service.durudex.local",
					Port: "8000",
					TLS: tls.PathConfig{
						Enable: true,
						CA:     "./certs/rootCA.pem",
						Cert:   "./certs/user.service.durudex.local-cert.pem",
						Key:    "./certs/user.service.durudex.local-key.pem",
					},
				},
				Database: config.DatabaseConfig{
					Postgres: config.PostgresConfig{
						MaxConns: 20,
						MinConns: 5,
						URL:      "postgres://localhost:1",
					},
				},
				Service: config.ServiceConfig{
					Code: grpc.ConnectionConfig{
						Addr: "code.service.durudex.local:8001",
						TLS: tls.PathConfig{
							Enable: true,
							CA:     "./certs/rootCA.pem",
							Cert:   "./certs/client-cert.pem",
							Key:    "./certs/client-key.pem",
						},
					},
				},
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environments configurations.
			setEnv(tt.args.env)

			// Creating a new config.
			got, err := config.New()
			if (err != nil) != tt.wantErr {
				t.Errorf("error creating a new config: %s", err.Error())
			}

			// Check for similarity of a config.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("error config are not similar")
			}
		})
	}
}
