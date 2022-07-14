/*
 * Copyright © 2022 Durudex
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

package auth_test

import (
	"testing"
	"time"

	"github.com/durudex/durudex-user-service/pkg/auth"
)

// Testing generating a new jwt access token.
func Test_GenerateAccessToken(t *testing.T) {
	// Testing args.
	type args struct {
		subject    string
		signingKey string
		ttl        time.Duration
	}

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				subject:    "1",
				signingKey: "secret-key",
				ttl:        time.Hour * 9999,
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate a new jwt access token.
			got, err := auth.GenerateAccessToken(tt.args.subject, tt.args.signingKey, tt.args.ttl)
			if (err != nil) != tt.wantErr {
				t.Errorf("error generating access token: %s", err.Error())
			}

			// Check access token is empty.
			if got == "" {
				t.Error("error access token is empty")
			}
		})
	}
}

// Testing generating a new refresh token.
func Test_GenerateRefreshToken(t *testing.T) {
	// Tests structures.
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "OK"},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate a new refresh token.
			got, err := auth.GenerateRefreshToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("error generating refresh token: %s", err.Error())
			}

			// Check refresh token is empty.
			if got == "" {
				t.Error("error refresh token is empty")
			}
		})
	}
}
