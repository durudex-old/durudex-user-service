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
	"time"

	"github.com/gofrs/uuid"
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
func (u User) Validate() error {
	switch {
	case !RxUsername.MatchString(u.Username):
		return &Error{Code: CodeInvalidArgument, Message: "Invalid Username"}
	case !RxPassword.MatchString(u.Password):
		return &Error{Code: CodeInvalidArgument, Message: "Invalid Password"}
	case !RxEmail.MatchString(u.Email):
		return &Error{Code: CodeInvalidArgument, Message: "Invalid Email"}
	}

	return nil
}
