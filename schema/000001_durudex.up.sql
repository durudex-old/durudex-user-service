/*
  Copyright © 2022 Durudex

  This file is part of Durudex: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  Durudex is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with Durudex. If not, see <https://www.gnu.org/licenses/>
*/

CREATE TABLE IF NOT EXISTS "user" (
  "id"         BIGSERIAL    NOT NULL UNIQUE,
  "username"   VARCHAR(40)  NOT NULL UNIQUE,
  "email"      VARCHAR(255) NOT NULL UNIQUE,
  "phone"      VARCHAR(15)           UNIQUE,
  "password"   VARCHAR(100) NOT NULL,
  "joined_in"  TIMESTAMP    NOT NULL,
  "last_join"  TIMESTAMP,
  "avatar_url" VARCHAR(255),
  "verified"   BOOLEAN      NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS "user_profile" (
  "id"       BIGSERIAL                                        NOT NULL UNIQUE,
  "user_id"  BIGINT REFERENCES "user"("id") ON DELETE CASCADE NOT NULL UNIQUE,
  "name"     VARCHAR(60),
  "sex"      SMALLINT                                         NOT NULL CHECK ("sex" IN (0, 1, 2, 9)),
  "status"   VARCHAR(100),
  "birthday" DATE                                             NULL NULL
);
