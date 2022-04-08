/*
 * Copyright © 2022 Durudex

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

package psql

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/durudex/durudex-user-service/internal/domain"

	"github.com/gofrs/uuid"
	"github.com/pashagolub/pgxmock"
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

	// Test bahavior.
	type mockBehavior func(args args, id uuid.UUID)

	// Creating a new repository.
	repos := NewUserRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		want         uuid.UUID
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{user: domain.User{
				Username: "example",
				Email:    "example@example.example",
				Password: "qwerty",
			}},
			want: uuid.UUID{},
			mockBehavior: func(args args, want uuid.UUID) {
				mock.ExpectQuery(`INSERT INTO "user"`).
					WithArgs(args.user.Username, args.user.Email, args.user.Password).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow(want))
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

// Testing getting user by credentials in postgres database.
func TestUserRepository_GetByCreds(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ username string }

	// Test bahavior.
	type mockBehavior func(args args, user domain.User)

	// Creating a new repository.
	repos := NewUserRepository(mock)

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
				ID:        uuid.UUID{},
				Username:  "example",
				Email:     "example@example.example",
				Password:  "qwerty123",
				JoinedIn:  time.Now(),
				LastVisit: time.Now(),
				Verified:  true,
				AvatarURL: nil,
			},
			mockBehavior: func(args args, user domain.User) {
				rows := mock.NewRows([]string{
					"id", "username", "email", "password", "joined_in", "last_visit",
					"verified", "avatar_url",
				}).AddRow(
					user.ID, user.Username, user.Email, user.Password, user.JoinedIn,
					user.LastVisit, user.Verified, user.AvatarURL)

				mock.ExpectQuery(`SELECT (.+) FROM "user"`).
					WithArgs(args.username).
					WillReturnRows(rows)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.want)

			// Getting user by credentials.
			got, err := repos.GetByCreds(context.Background(), tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("error getting user by credentials: %s", err.Error())
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

	// Test bahavior.
	type mockBehavior func(args args)

	// Creating a new repository.
	repos := NewUserRepository(mock)

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
				mock.ExpectExec(`UPDATE "user"`).
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
