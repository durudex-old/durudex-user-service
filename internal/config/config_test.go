/*
	Copyright © 2021-2022 Durudex

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
	"os"
	"reflect"
	"testing"
	"time"
)

// Test initialize config.
func TestConfig_Init(t *testing.T) {
	// Environment configurations.
	type env struct{ configPath, postgresURL, redisURL string }

	// Testing args.
	type args struct{ env env }

	// Set environments configurations.
	setEnv := func(env env) {
		os.Setenv("CONFIG_PATH", env.configPath)
		os.Setenv("POSTGRES_URL", env.postgresURL)
		os.Setenv("REDIS_URL", env.redisURL)
	}

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "OK",
			args: args{env: env{
				configPath:  "fixtures/main",
				postgresURL: "postgres://localhost:1",
				redisURL:    "redis://user.redis.durudex.local:6379",
			}},
			want: &Config{
				GRPC: GRPCConfig{
					Host: "user.service.durudex.local",
					Port: "8004",
					TLS: TLSConfig{
						Enable: true,
						CACert: "./certs/rootCA.pem",
						Cert:   "./certs/sample.service.durudex.local-cert.pem",
						Key:    "./certs/sample.service.durudex.local-key.pem",
					},
				},
				Database: DatabaseConfig{
					Postgres: PostgresConfig{
						MaxConns: 20,
						MinConns: 5,
						URL:      "postgres://localhost:1",
					},
					Redis: RedisConfig{URL: "redis://user.redis.durudex.local:6379"},
				},
				Password: PasswordConfig{Cost: 14},
				Code:     CodeConfig{TTL: time.Minute * 15},
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environments configurations.
			setEnv(tt.args.env)

			// Initialize connfig.
			got, err := Init()
			if (err != nil) != tt.wantErr {
				t.Errorf("error initialize config: %s", err.Error())
			}

			// Check for similarity of a config.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("error config are not similar")
			}
		})
	}
}
