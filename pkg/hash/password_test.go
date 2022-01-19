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

package hash

import "testing"

const (
	password = "1234567890"
	cost     = 14
)

// Testing generating a new password hash.
func TestHash(t *testing.T) {
	// Creating a new password manager.
	manager := NewPassword(cost)

	// Generating a new password hash.
	hash, err := manager.Hash(password)
	if err != nil {
		t.Errorf("error hashing password: %s", err.Error())
	}

	// Check for availability.
	if hash == "" {
		t.Error("password hash is empty")
	}
}

// Testing checking password hash.
func TestCheck(t *testing.T) {
	// Creating a new password manager.
	manager := NewPassword(cost)

	// Generating a new password hash.
	hash, err := manager.Hash(password)
	if err != nil {
		t.Errorf("error hashing password: %s", err.Error())
	}

	// Check password hash.
	status := manager.Check(hash, password)
	if !status {
		t.Error("password do not match")
	}
}
