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

package domain

import (
	"net/mail"
	"regexp"

	"github.com/durudex/go-shared/status"

	"github.com/segmentio/ksuid"
)

// The credential key is required to retrieve the user's private data from the database.
type CredentialKey string

// Getting the credential key in string.
func (k CredentialKey) String() string { return string(k) }

const (
	// Credential key used to select by username.
	UsernameCredential CredentialKey = "username"
	// Credential key used to select by email address.
	EmailCredential CredentialKey = "email"

	// A regular expression pattern for validating a username.
	UsernamePattern string = "^[a-zA-Z0-9-_.]{3,40}$"
	// A regular expression pattern for validating a password.
	PasswordPattern string = "^[a-zA-Z0-9]{64}$"
)

var (
	// A regular expression for validating a username.
	RegularUsername = regexp.MustCompile(UsernamePattern)
	// A regular expression for validating a password.
	RegularPassword = regexp.MustCompile(PasswordPattern)
)

// Durudex user structure.
type User struct {
	// This is a unique user ID in the Durudex network that has the KSUID format.
	ID ksuid.KSUID

	// The unique username on the Durudex network.
	Username string

	// The unique email address of the user's account on the Durudex network.
	Email string

	// Server-hashed user account password.
	PasswordHash string

	// The hash epoch at which the user account password was hashed.
	PasswordEpoch int

	// The status of a user who has passed an identity verification check.
	Verified bool

	// User account avatar url on the Durudex network.
	AvatarURL *string
}

// Input for creating a new user.
type CreateUserInput struct {
	// This is a unique user ID in the Durudex network that has the KSUID format.
	ID ksuid.KSUID

	// The unique username on the Durudex network.
	Username string

	// The unique email address of the user account on the Durudex network.
	Email string

	// Client-hashed user account password.
	PasswordHash string

	// The hash epoch at which the user account password was hashed.
	PasswordEpoch int
}

// Validation of user data according to Durudex standards.
func (i CreateUserInput) Validate() error {
	// Validation of e-mail address.
	if _, err := mail.ParseAddress(i.Email); err != nil {
		return &status.Error{
			Code: status.CodeInvalidArgument, Message: "Invalid `email` field",
		}
	}

	switch {
	case !RegularUsername.MatchString(i.Username):
		return &status.Error{
			Code: status.CodeInvalidArgument, Message: "Invalid `username` field",
		}
	case !RegularPassword.MatchString(i.PasswordHash):
		return &status.Error{
			Code: status.CodeInvalidArgument, Message: "Invalid password hash",
		}
	default:
		return nil
	}
}
