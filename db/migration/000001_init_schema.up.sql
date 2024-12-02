CREATE TABLE "wallets" (
                           "id" bigserial PRIMARY KEY,
                           "type" varchar NOT NULL,
                           "full_name" varchar NOT NULL,
                           "document" varchar UNIQUE NOT NULL,
                           "email" varchar UNIQUE NOT NULL,
                           "password" varchar NOT NULL,
                           "balance" bigint NOT NULL DEFAULT 0,
                           "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "entries" (
                           "id" bigserial PRIMARY KEY,
                           "wallet_id" bigint NOT NULL,
                           "amount" bigint NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "transfers" (
                             "id" bigserial PRIMARY KEY,
                             "from_wallet" bigint NOT NULL,
                             "to_wallet" bigint NOT NULL,
                             "amount" bigint NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "wallets" ("document");

CREATE INDEX ON "wallets" ("email");

CREATE INDEX ON "entries" ("wallet_id");

CREATE INDEX ON "transfers" ("from_wallet");

CREATE INDEX ON "transfers" ("to_wallet");

CREATE INDEX ON "transfers" ("from_wallet", "to_wallet");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "entries" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_wallet") REFERENCES "wallets" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_wallet") REFERENCES "wallets" ("id");
