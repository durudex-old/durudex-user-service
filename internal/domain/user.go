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
)

// User model.
type User struct {
	ID        uint64
	Username  string
	Email     string
	Password  string
	JoinedIn  time.Time
	LastVisit time.Time
	Verified  bool
	AvatarURL *string
}

// Validate user.
func (u *User) Validate() error {
	switch {
	case !rxUsername.MatchString(u.Username):
		return errors.New("error username is incorrect")
	case !rxPassword.MatchString(u.Password):
		return errors.New("error password is incorrect")
	case !rxEmail.MatchString(u.Email):
		return errors.New("error email is incorrect")
	}

	return nil
}
