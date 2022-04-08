# Copyright © 2021-2022 Durudex

# This file is part of Durudex: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as
# published by the Free Software Foundation, either version 3 of the
# License, or (at your option) any later version.

# Durudex is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.

# You should have received a copy of the GNU Affero General Public License
# along with Durudex. If not, see <https://www.gnu.org/licenses/>.

POSTGRES_URL=postgresql://admin:qwerty@user.postgres.durudex.local:5433/durudex

.PHONY: download
download:
	go mod download

.PHONY: build
build:
	go mod download && CGO_ENABLE=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

.PHONY: run
run: build
	docker-compose up --remove-orphans app

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test: lint
	go test -v ./...

.PHONY: migrate-up
migrate-up:
	migrate -path ./schema/migrations -database '$(POSTGRES_URL)?sslmode=disable' up

.PHONY: migrate-down
migrate-down:
	migrate -path ./schema/migrations -database '$(POSTGRES_URL)?sslmode=disable' down

.PHONY: protoc
protoc:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		internal/delivery/grpc/pb/*.proto

.PHONY: protoc-types
protoc-types:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		internal/delivery/grpc/pb/types/*.proto

.DEFAULT_GOAL := run
