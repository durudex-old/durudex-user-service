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

// Testing creating a new user session in postgres database.
func TestSessionRepository_Create(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ session domain.Session }

	// Test behavior.
	type mockBehavior func(args args)

	// Creating a new repository.
	repos := postgres.NewSessionRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{domain.Session{
				Id:           ksuid.New(),
				UserId:       ksuid.New(),
				RefreshToken: "qwerty",
				Ip:           "0.0.0.0",
				ExpiresIn:    time.Now(),
			}},
			mockBehavior: func(args args) {
				query := fmt.Sprintf(`INSERT INTO "%s"`, postgres.SessionTable)
				mock.ExpectExec(query).
					WithArgs(args.session.Id, args.session.UserId, args.session.RefreshToken,
						args.session.Ip, args.session.ExpiresIn).
					WillReturnResult(pgxmock.NewResult("", 1))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			// Creating a new user session in postgres database.
			err := repos.Create(context.Background(), tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("error creating a new user session: %s", err.Error())
			}
		})
	}
}

// Testing getting user id by refresh token in postgres database.
func TestSessionRepository_GetUserId(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ refreshToken, ip string }

	// Test behavior.
	type mockBehavior func(args args, id ksuid.KSUID)

	// Creating a new repository.
	repos := postgres.NewSessionRepository(mock)

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
			args: args{refreshToken: "qwerty", ip: "0.0.0.0"},
			want: ksuid.New(),
			mockBehavior: func(args args, id ksuid.KSUID) {
				query := fmt.Sprintf(`SELECT (.+) FROM "%s"`, postgres.SessionTable)
				mock.ExpectQuery(query).
					WithArgs(args.refreshToken, args.ip).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow(id.String()))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.want)

			// Get user id by refresh token in postgres database.
			got, err := repos.GetUserId(context.Background(), tt.args.refreshToken, tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("error getting user id by refresh token: %s", err.Error())
			}

			// Check for similarity of user id.
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error user id are not similar")
			}
		})
	}
}

// Testing deleting a user session in postgres database.
func TestSessionRepository_Delete(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ refreshToken, ip string }

	// Test behavior.
	type mockBehavior func(args args)

	// Creating a new repository.
	repos := postgres.NewSessionRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name:    "OK",
			args:    args{refreshToken: "qwerty", ip: "0.0.0.0"},
			wantErr: false,
			mockBehavior: func(args args) {
				query := fmt.Sprintf(`DELETE FROM "%s"`, postgres.SessionTable)
				mock.ExpectExec(query).
					WithArgs(args.refreshToken, args.ip).
					WillReturnResult(pgxmock.NewResult("", 1))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			// Deleting a user session in postgres database.
			err := repos.Delete(context.Background(), tt.args.refreshToken, tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("error deleting user session: %s", err.Error())
			}
		})
	}
}
