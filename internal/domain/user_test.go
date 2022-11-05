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

package domain_test

import (
	"testing"

	"github.com/durudex/durudex-user-service/internal/domain"
)

// Testing validation of user data according to Durudex standards.
func TestCreateUserInput_Validate(t *testing.T) {
	// Testing arguments.
	type args struct{ username, email, passwordHash string }

	// Testing structures.
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				username:     "Example",
				email:        "example@durudex.com",
				passwordHash: "91b9b4ddda35be0338407fbaa76bb6adfe2dba8ad6719fe0ebae006c297b529f",
			},
			wantErr: false,
		},
		{
			name: "Email Not Correct",
			args: args{
				username:     "Test",
				email:        "example.example",
				passwordHash: "91b9b4ddda35be0338407fbaa76bb6adfe2dba8ad6719fe0ebae006c297b529f",
			},
			wantErr: true,
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Creating a new input for creating a new user.
			user := domain.CreateUserInput{
				Username:     tt.args.username,
				Email:        tt.args.email,
				PasswordHash: tt.args.passwordHash,
			}

			// Validation of user data according to Durudex standards.
			if err := user.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("error validation CreateUserInput: %s", err.Error())
			}
		})
	}
}
