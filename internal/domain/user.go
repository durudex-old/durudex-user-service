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

package domain

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

var (
	ErrPasswordIsIncorrect = errors.New("error password is incorrect")
	ErrUsernameIsIncorrect = errors.New("error username is incorrect")
	ErrEmailIsIncorrect    = errors.New("error email is incorrect")
)

// User model.
type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	LastVisit time.Time
	Verified  bool
	AvatarURL *string
}

// Validate user.
func (u *User) Validate() error {
	switch {
	case !RxUsername.MatchString(u.Username):
		return ErrUsernameIsIncorrect
	case !RxPassword.MatchString(u.Password):
		return ErrPasswordIsIncorrect
	case !RxEmail.MatchString(u.Email):
		return ErrEmailIsIncorrect
	}

	return nil
}
