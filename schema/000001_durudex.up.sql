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

CREATE TABLE IF NOT EXISTS users (
  id             CHAR(27)     NOT NULL,
  username       VARCHAR(40)  NOT NULL,
  email          VARCHAR(255) NOT NULL,
  password_hash  VARCHAR(64)  NOT NULL,
  password_epoch SMALLINT     NOT NULL
  verified       BOOLEAN      NOT NULL DEFAULT false,
  avatar_url     VARCHAR(255),
  CONSTRAINT users_pkey PRIMARY KEY (id DESC)
);

CREATE UNIQUE INDEX IF NOT EXISTS users_username_key (username DESC);
CREATE UNIQUE INDEX IF NOT EXISTS users_email_key (email DESC);
