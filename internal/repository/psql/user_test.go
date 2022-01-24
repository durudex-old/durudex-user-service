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

	"github.com/pashagolub/pgxmock"
)

// Testing creating a new user in postgres datatabe.
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
	type mockBehavior func(args args, id uint64)

	// Creating a new repository.
	repos := NewUserRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		want         uint64
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{user: domain.User{
				Username: "Test",
				Email:    "example@example.example",
				Password: "superpassword",
			}},
			want: 1,
			mockBehavior: func(args args, want uint64) {
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

			// Creating a new user in postgres datatabe.
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

// Testing getting user by credentials for postgres database.
func TestUserRepository_GetByCreds(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct {
		username string
		password string
	}

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
			args: args{
				username: "Test",
				password: "superpassword",
			},
			want: domain.User{
				ID:        1,
				Username:  "Test",
				Email:     "example@example.example",
				JoinedIn:  time.Now(),
				LastVisit: time.Now(),
				Verified:  true,
				AvatarURL: nil,
			},
			mockBehavior: func(args args, user domain.User) {
				rows := mock.NewRows([]string{
					"id", "username", "email", "joined_in", "last_visit", "verified",
					"avatar_url",
				}).AddRow(
					user.ID, user.Username, user.Email, user.JoinedIn, user.LastVisit,
					user.Verified, user.AvatarURL)

				mock.ExpectQuery(`SELECT (.+) FROM "user"`).
					WithArgs(args.username, args.password).
					WillReturnRows(rows)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.want)

			// Getting user by credentials.
			got, err := repos.GetByCreds(context.Background(), tt.args.username, tt.args.password)
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