CREATE TABLE "users"
(
    "username"        varchar primary key,
    "hashed_pass"     varchar        not null,
    "full_name"       varchar        not null,
    "emain"           varchar unique not null,
    "pass_changed_at" timestamptz    not null default '0001-01-01 00:00:00Z',
    "created_at"      timestamptz    not null default NOW()
);

ALTer table "accounts"
    add foreign key ("owner") references "users" ("username");

-- create unique index on "accounts" ("owner", "currency");
alter table "accounts"
    add constraint "owner_currency_key" unique ("owner", "currency");