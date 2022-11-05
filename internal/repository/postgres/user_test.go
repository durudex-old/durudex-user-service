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

package postgres_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/internal/repository/postgres"

	"github.com/pashagolub/pgxmock/v2"
	"github.com/segmentio/ksuid"
)

// Testing creating a new user in the database.
func TestUserRepository_Create(t *testing.T) {
	// Creating a new mock pool connection.
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error creating a new mock pool connection: %s", err.Error())
	}
	defer mock.Close()

	// Testing args.
	type args struct{ input domain.CreateUserInput }

	// Test behavior.
	type mockBehavior func(args args)

	// Creating a new repository.
	repos := postgres.NewUserRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{input: domain.CreateUserInput{
				ID:            ksuid.New(),
				Username:      "example",
				Email:         "example@durudex.com",
				PasswordHash:  "91b9b4ddda35be0338407fbaa76bb6adfe2dba8ad6719fe0ebae006c297b529f",
				PasswordEpoch: 1,
			}},
			mockBehavior: func(args args) {
				mock.ExpectExec("INSERT INTO users").
					WithArgs(args.input.ID, args.input.Username, args.input.Email, args.input.PasswordHash,
						args.input.PasswordEpoch).
					WillReturnResult(pgxmock.NewResult("", 1))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			// Creating a new user in the database.
			if err := repos.Create(context.Background(), tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("error creating a new user: %s", err.Error())
			}
		})
	}
}

// Testing getting a user from the database by his ID.
func TestUserRepository_Get(t *testing.T) {
	// Creating a new mock pool connection.
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error creating a new mock pool connection: %s", err.Error())
	}
	defer mock.Close()

	// Testing args.
	type args struct{ uid ksuid.KSUID }

	// Test behavior.
	type mockBehavior func(args args, user domain.User)

	// Creating a new repository.
	repos := postgres.NewUserRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		want         domain.User
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{uid: ksuid.New()},
			want: domain.User{
				Username:  "example",
				Verified:  true,
				AvatarURL: nil,
			},
			mockBehavior: func(args args, user domain.User) {
				rows := mock.NewRows([]string{
					"username", "verified", "avatar_url",
				}).AddRow(user.Username, user.Verified, user.AvatarURL)

				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs(args.uid).
					WillReturnRows(rows)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.want)

			// Getting a user from the database by his ID.
			got, err := repos.Get(context.Background(), tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("error getting user: %s", err.Error())
			}

			// Check for similarity of user.
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error user are not similar")
			}
		})
	}
}

// Testing getting a user from the database using credentials
func TestUserRepository_GetByCreds(t *testing.T) {
	// Creating a new mock pool connection.
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error creating a new mock pool connection: %s", err.Error())
	}
	defer mock.Close()

	// Testing args.
	type args struct {
		key   domain.CredentialKey
		value string
	}

	// Test behavior.
	type mockBehavior func(args args, user domain.User)

	// Creating a new repository.
	repos := postgres.NewUserRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		want         domain.User
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{key: domain.UsernameCredential, value: "example"},
			want: domain.User{
				ID:            ksuid.New(),
				Username:      "example",
				Email:         "example@durudex.com",
				PasswordHash:  "91b9b4ddda35be0338407fbaa76bb6adfe2dba8ad6719fe0ebae006c297b529f",
				PasswordEpoch: 1,
				Verified:      true,
				AvatarURL:     nil,
			},
			mockBehavior: func(args args, user domain.User) {
				rows := mock.NewRows([]string{
					"id", "username", "email", "password_hash",
					"password_epoch", "verified", "avatar_url",
				}).AddRow(user.ID, user.Username, user.Email, user.PasswordHash,
					user.PasswordEpoch, user.Verified, user.AvatarURL)

				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs(args.value).
					WillReturnRows(rows)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.want)

			// Getting a user from the database using credentials.
			got, err := repos.GetByCreds(context.Background(), tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("error getting user by credential: %s", err.Error())
			}

			// Check for similarity of user.
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error user are not similar")
			}
		})
	}
}
