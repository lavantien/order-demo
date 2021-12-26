CREATE TABLE
    "users" (
        "username" VARCHAR PRIMARY KEY,
        "hashed_password" VARCHAR NOT NULL,
        "full_name" VARCHAR NOT NULL,
        "email" VARCHAR UNIQUE NOT NULL,
        "password_change_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );

CREATE TABLE
    "orders" (
        "id" bigserial PRIMARY KEY,
        "owner" VARCHAR NOT NULL,
        "product_id" BIGINT NOT NULL,
        "quantity" BIGINT NOT NULL,
        "price" BIGINT NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );

CREATE TABLE
    "products" (
        "id" bigserial PRIMARY KEY,
        "name" VARCHAR NOT NULL,
        "cost" BIGINT NOT NULL,
        "quantity" BIGINT NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );

ALTER TABLE
    "orders"
ADD
    FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE
    "orders"
ADD
    FOREIGN KEY ("product_id") REFERENCES "products" ("id");

CREATE INDEX
ON "users" ("hashed_password");

CREATE INDEX
ON "users" ("full_name");

CREATE INDEX
ON "users" ("password_change_at");

CREATE INDEX
ON "orders" ("owner");

CREATE INDEX
ON "orders" ("product_id");

CREATE INDEX
ON "orders" ("quantity");

CREATE INDEX
ON "orders" ("price");

CREATE INDEX
ON "products" ("name");

CREATE INDEX
ON "products" ("cost");

CREATE INDEX
ON "products" ("quantity");

COMMENT
ON COLUMN "orders"."quantity" IS 'must be positive';

COMMENT
ON COLUMN "orders"."price" IS 'must be positive';

COMMENT
ON COLUMN "products"."cost" IS 'must be positive';

COMMENT
ON COLUMN "products"."quantity" IS 'must be positive';

-- Generate admin
INSERT INTO
    users (
        username,
        hashed_password,
        full_name,
        email,
        password_change_at,
        created_at
    )
VALUES (
        'admin',
        '$2a$10$VxkKRxRSov1e2LzNXc1aden5kkDAJEM5RF5n60HauC/zLpFhx/jfe',
        'Admin',
        'admin@email.com',
        '0001-01-01 07:00:00.000',
        '2021-12-26 22:22:49.644'
    );
