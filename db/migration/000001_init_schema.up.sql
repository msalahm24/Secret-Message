CREATE TABLE "users" (
  "id" bigserial,
  "user_name" varchar PRIMARY KEY,
  "phone_number" varchar NOT NULL,
  "location" varchar,
  "type_of_device" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "messages" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "to_phone_number" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "content" text NOT NULL
);

CREATE TABLE "phone_numbers" (
  "id" bigserial PRIMARY KEY,
  "phone_number" varchar UNIQUE NOT NULL,
  "location" varchar,
  "user_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("user_name");

CREATE INDEX ON "users" ("phone_number");

CREATE INDEX ON "messages" ("user_id");

CREATE INDEX ON "messages" ("to_phone_number");

CREATE INDEX ON "phone_numbers" ("user_id");

CREATE INDEX ON "phone_numbers" ("phone_number");

ALTER TABLE "messages" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_name");

ALTER TABLE "phone_numbers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_name");
