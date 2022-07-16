/*
 * Copyright © 2021-2022 Durudex
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

package service

import (
	"github.com/durudex/durudex-user-service/internal/config"
	"github.com/durudex/durudex-user-service/internal/repository"
	v1 "github.com/durudex/durudex-user-service/pkg/pb/durudex/v1"
)

// Service structure.
type Service struct {
	User
	Auth
	Code
}

// Creating a new service.
func NewService(repos *repository.Repository, config *config.Config, email v1.EmailServiceClient) *Service {
	codeService := NewCodeService(repos.Redis, email, &config.Code)
	userService := NewUserService(repos.Postgres.User, codeService, &config.Password)

	return &Service{
		User: userService,
		Auth: &AuthService{
			user:    userService,
			code:    codeService,
			email:   email,
			session: repos.Postgres.Session,
			cfg:     &config.Auth,
		},
		Code: codeService,
	}
}
