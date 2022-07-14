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

package rand

import (
	"crypto/rand"
	"math/big"
)

// Code manager interface.
type Code interface {
	Generate() (uint64, error)
}

// Generating a random uint64 code.
func Generate(maxLength, minLength int64) (uint64, error) {
	// Creating big int.
	bg := big.NewInt(maxLength - minLength)

	// Generating random code.
	code, err := rand.Int(rand.Reader, bg)
	if err != nil {
		return 0, err
	}

	return code.Uint64() + uint64(minLength), nil
}
