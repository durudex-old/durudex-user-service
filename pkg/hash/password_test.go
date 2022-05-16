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

package hash_test

import (
	"testing"

	"github.com/durudex/durudex-user-service/pkg/hash"
)

// Testing generating a new password hash.
func Test_Hash(t *testing.T) {
	// Testing args.
	type args struct {
		cost     int
		password string
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
				cost:     14,
				password: "1234567890",
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := hash.Hash(tt.args.password, tt.args.cost)
			if (err != nil) != tt.wantErr {
				t.Errorf("error hashing password: %s", err.Error())
			}

			// Check password hash.
			if got == "" {
				t.Errorf("error password hash is empty")
			}
		})
	}
}

// Testing checking password hash.
func Test_Check(t *testing.T) {
	// Testing args.
	type args struct {
		cost     int
		password string
		hash     string
	}

	// Tests structures.
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "OK",
			args: args{
				cost:     14,
				password: "1234567890",
				hash:     "$2a$14$SnPQXou3EwjQHDgKb0/b.uKgwD2PRNVVV9m5s4RxE7Zu9v.zL1bSq",
			},
			want: true,
		},
		{
			name: "Password Not Correct",
			args: args{
				cost:     14,
				password: "ne1234567890",
				hash:     "$2a$14$SnPQXou3EwjQHDgKb0/b.uKgwD2PRNVVV9m5s4RxE7Zu9v.zL1bSq",
			},
			want: false,
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check password hash.
			got := hash.Check(tt.args.hash, tt.args.password)
			if got != tt.want {
				t.Error("password are not similar")
			}
		})
	}
}
