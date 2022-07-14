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

package rand_test

import (
	"testing"

	"github.com/durudex/durudex-user-service/pkg/crypto/rand"
)

// Testing generating a random uint64 code.
func Test_Generate(t *testing.T) {
	// Testing args.
	type args struct{ maxLength, minLength int64 }

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{maxLength: 999999, minLength: 100000},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generating a random uint64 code.
			_, err := rand.Generate(tt.args.maxLength, tt.args.minLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("error generating code: %s", err.Error())
			}
		})
	}
}
