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

import "golang.org/x/crypto/bcrypt"

// Password hash manager.
type PasswordManager struct{ Cost int }

// Creating a new password hash manager.
func NewPassword(cost int) *PasswordManager {
	return &PasswordManager{Cost: cost}
}

// Generating a new password hash.
func (m *PasswordManager) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), m.Cost)
	return string(bytes), err
}

// Check password hash.
func (m *PasswordManager) Check(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
