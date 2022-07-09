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
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/durudex/durudex-user-service/internal/domain"
	"github.com/durudex/durudex-user-service/internal/repository/postgres"

	"github.com/pashagolub/pgxmock"
	"github.com/segmentio/ksuid"
)

// Testing creating a new user in postgres database.
func TestUserRepository_Create(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ user domain.User }

	// Test behavior.
	type mockBehavior func(args args, id ksuid.KSUID)

	// Creating a new repository.
	repos := postgres.NewUserRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		want         ksuid.KSUID
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{user: domain.User{
				Username: "example",
				Email:    "example@durudex.com",
				Password: "qwerty",
			}},
			want: ksuid.New(),
			mockBehavior: func(args args, want ksuid.KSUID) {
				mock.ExpectQuery(fmt.Sprintf(`INSERT INTO "%s"`, postgres.UserTable)).
					WithArgs(args.user.Username, args.user.Email, args.user.Password).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow(want.String()))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.want)

			// Creating a new user in postgres database.
			got, err := repos.Create(context.Background(), tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("error creating user: %s", err.Error())
			}

			// Check for similarity of id.
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error id are not similar")
			}
		})
	}
}

// Testing getting user by id in postgres database.
func TestUserRepository_GetByID(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ id ksuid.KSUID }

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
			args: args{id: ksuid.New()},
			want: domain.User{
				Username:  "example",
				LastVisit: time.Now(),
				Verified:  true,
				AvatarUrl: nil,
			},
			mockBehavior: func(args args, user domain.User) {
				rows := mock.NewRows([]string{
					"username", "last_visit", "verified", "avatar_url",
				}).AddRow(user.Username, user.LastVisit, user.Verified, user.AvatarUrl)

				mock.ExpectQuery(fmt.Sprintf(`SELECT (.+) FROM "%s"`, postgres.UserTable)).
					WithArgs(args.id.String()).
					WillReturnRows(rows)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.want)

			// Getting user by id.
			got, err := repos.GetByID(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error getting user by id: %s", err.Error())
			}

			// Check for similarity of user.
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error user are not similar")
			}
		})
	}
}

// Testing getting user by username in postgres database.
func TestUserRepository_GetByUsername(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ username string }

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
			args: args{username: "example"},
			want: domain.User{
				Id:        ksuid.New(),
				Email:     "example@example.example",
				Password:  "qwerty123",
				LastVisit: time.Now(),
				Verified:  true,
				AvatarUrl: nil,
			},
			mockBehavior: func(args args, user domain.User) {
				rows := mock.NewRows([]string{
					"id", "email", "password", "last_visit", "verified", "avatar_url",
				}).AddRow(user.Id.String(), user.Email, user.Password, user.LastVisit,
					user.Verified, user.AvatarUrl)

				mock.ExpectQuery(fmt.Sprintf(`SELECT (.+) FROM "%s"`, postgres.UserTable)).
					WithArgs(args.username).
					WillReturnRows(rows)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.want)

			// Getting user by username.
			got, err := repos.GetByUsername(context.Background(), tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("error getting user by username: %s", err.Error())
			}

			// Check for similarity of user.
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error user are not similar")
			}
		})
	}
}

// Testing forgot password in postgres database.
func TestUserRepository_ForgotPassword(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ email, password string }

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
			args: args{email: "example@example.example", password: "qwerty"},
			mockBehavior: func(args args) {
				mock.ExpectExec(fmt.Sprintf(`UPDATE "%s"`, postgres.UserTable)).
					WithArgs(args.password, args.email).
					WillReturnResult(pgxmock.NewResult("", 1))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			// Forgot password in postgres database.
			err := repos.ForgotPassword(context.Background(), tt.args.password, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("error forgot user password: %s", err.Error())
			}
		})
	}
}

// Testing update user in postgres database.
func TestUserRepository_UpdateAvatar(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct {
		avatarUrl string
		id        ksuid.KSUID
	}

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
			args: args{
				avatarUrl: "https://cdn.durudex.com/avatar/0ujzPyRiIAffKhBux4PvQdDqMHY/user.png",
				id:        ksuid.New(),
			},
			mockBehavior: func(args args) {
				mock.ExpectExec(fmt.Sprintf(`UPDATE "%s"`, postgres.UserTable)).
					WithArgs(args.avatarUrl, args.id.String()).
					WillReturnResult(pgxmock.NewResult("", 1))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			// Update user avatar in postgres database.
			err := repos.UpdateAvatar(context.Background(), tt.args.avatarUrl, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error updating user avatar: %s", err.Error())
			}
		})
	}
}
