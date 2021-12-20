CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
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

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("hashed_password");

CREATE INDEX ON "orders" ("user_id");

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
