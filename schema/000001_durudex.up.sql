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

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "user" (
  "id"         UUID         NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "username"   VARCHAR(40)  NOT NULL UNIQUE,
  "email"      VARCHAR(255) NOT NULL UNIQUE,
  "password"   VARCHAR(100) NOT NULL,
  "created_at" TIMESTAMP    NOT NULL DEFAULT now(),
  "last_visit" TIMESTAMP    NOT NULL DEFAULT now(),
  "verified"   BOOLEAN      NOT NULL DEFAULT false,
  "avatar_url" VARCHAR(255)
);
