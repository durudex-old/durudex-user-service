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

package service

import "context"

// Code service interface.
type Code interface {
	CreateVerifyEmailCode(ctx context.Context, email string) error
	VerifyEmailCode(ctx context.Context, email string, code uint64) (bool, error)
}

// Code service structure.
type CodeService struct{}

// Creating a new code service.
func NewCodeService() *CodeService {
	return &CodeService{}
}

func (s *CodeService) CreateVerifyEmailCode(ctx context.Context, email string) error {
	return nil
}

func (s *CodeService) VerifyEmailCode(ctx context.Context, email string, code uint64) (bool, error) {
	return false, nil
}
