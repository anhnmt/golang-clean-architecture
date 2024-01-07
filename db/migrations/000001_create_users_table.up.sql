START TRANSACTION;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id         uuid        NOT NULL DEFAULT (uuid_generate_v4()),
    name       varchar     NOT NULL,
    email      varchar     NOT NULL,
    password   varchar     NOT NULL,
    created_at timestamptz not null DEFAULT (now()),
    updated_at timestamptz not null DEFAULT (now()),
    CONSTRAINT pk_users PRIMARY KEY (id)
);

create unique index user_id
    on users (id);

create index user_created_at
    on users (created_at);

COMMIT;