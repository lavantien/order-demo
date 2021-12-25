CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_change_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "product_id" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "cost" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "orders" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

CREATE INDEX ON "users" ("hashed_password");

CREATE INDEX ON "users" ("full_name");

CREATE INDEX ON "users" ("password_change_at");

CREATE INDEX ON "orders" ("owner");

CREATE INDEX ON "orders" ("product_id");

CREATE INDEX ON "orders" ("quantity");

CREATE INDEX ON "orders" ("price");

CREATE INDEX ON "products" ("name");

CREATE INDEX ON "products" ("cost");

CREATE INDEX ON "products" ("quantity");

COMMENT ON COLUMN "orders"."quantity" IS 'must be positive';

COMMENT ON COLUMN "orders"."price" IS 'must be positive';

COMMENT ON COLUMN "products"."cost" IS 'must be positive';

COMMENT ON COLUMN "products"."quantity" IS 'must be positive';
