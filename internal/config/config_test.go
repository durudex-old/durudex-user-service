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
)

// Test initialize config.
func TestConfig_Init(t *testing.T) {
	// Environment configurations.
	type env struct{ configPath, postgresURL string }

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
		want    *Config
		wantErr bool
	}{
		{
			name: "OK",
			args: args{env: env{
				configPath:  "fixtures/main",
				postgresURL: "postgres://localhost:1",
			}},
			want: &Config{
				Server: ServerConfig{
					Host: defaultServerHost,
					Port: defaultServerPort,
					TLS: TLSConfig{
						Enable: defaultTLSEnable,
						CACert: defaultTLSCACert,
						Cert:   defaultTLSCert,
						Key:    defaultTLSKey,
					},
				},
				Database: DatabaseConfig{
					Postgres: PostgresConfig{
						MaxConns: defaultPostgresMaxConns,
						MinConns: defaultPostgresMinConns,
						URL:      "postgres://localhost:1",
					},
				},
				Hash: HashConfig{Password: PasswordConfig{Cost: defaultPasswordCost}},
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
