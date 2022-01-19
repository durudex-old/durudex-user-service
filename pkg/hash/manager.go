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

// Password hash interface.
type Password interface {
	Hash(password string) (string, error)
	Check(hash, password string) bool
}

// Hash manager structure.
type Hash struct{ Password }

// Creating a new hash manager.
func NewHash(cost int) *Hash {
	return &Hash{Password: NewPassword(cost)}
}
