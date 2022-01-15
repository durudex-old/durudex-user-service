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
	"reflect"
	"testing"
)

// Test initialize config.
func TestInit(t *testing.T) {
	// Testing args.
	type args struct{ path string }

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		wand    *Config
		wantErr bool
	}{
		{
			name: "test config",
			args: args{path: "fixtures/main"},
			wand: &Config{
				Server: ServerConfig{
					Host: defaultServerHost,
					Port: defaultServerPort,
					TLS:  defaultServerTLS,
				},
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize connfig.
			got, err := Init(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("error initialize config: %s", err.Error())
			}

			// Check for similarity of a config.
			if !reflect.DeepEqual(got, tt.wand) {
				t.Errorf("error config are not similar: %s", err.Error())
			}
		})
	}
}
